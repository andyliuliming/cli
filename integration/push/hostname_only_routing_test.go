package push

import (
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("push with hostname", func() {
	var (
		appName string
		route   string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()
	})

	Context("when the default domain is a HTTP domain", func() {
		Context("when no host is provided / host defaults to app name", func() {
			BeforeEach(func() {
				route = fmt.Sprintf("%s.%s", strings.ToLower(appName), helpers.DefaultSharedDomain())
			})

			Context("when the default route does not exist", func() {
				It("creates and maps the route", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--no-start")
						Eventually(session).Should(Say("routes:"))
						Eventually(session).Should(Say("(?i)\\+\\s+%s", route))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Say("routes:\\s+%s", route))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("the default route exists and is unmapped", func() {
				BeforeEach(func() {
					Eventually(helpers.CF("create-route", space, helpers.DefaultSharedDomain(), "-n", strings.ToLower(appName))).Should(Exit(0))
				})

				It("maps the route", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--no-start")
						Eventually(session).Should(Say("routes:"))
						Eventually(session).Should(Say("(?i)\\+\\s+%s", route))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Say("routes:\\s+%s", route))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the default route is mapped to the application", func() {
				BeforeEach(func() {
					helpers.WithHelloWorldApp(func(dir string) {
						Eventually(helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--no-start")).Should(Exit(0))
					})
				})

				It("does nothing", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--no-start")
						Eventually(session).Should(Say("routes:"))
						Eventually(session).Should(Say("(?i)\\s+%s", route))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Say("routes:\\s+%s", route))
					Eventually(session).Should(Exit(0))
				})
			})
		})

		Context("when the host is provided", func() {
			var (
				hostname string
			)

			BeforeEach(func() {
				hostname = strings.ToLower(helpers.RandomName())
				route = fmt.Sprintf("%s.%s", hostname, helpers.DefaultSharedDomain())
			})

			Context("when the default route does not exist", func() {
				Context("when no other routes exist", func() {
					It("creates and maps the route, and prevents mapping of the default route", func() {
						helpers.WithHelloWorldApp(func(dir string) {
							session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--hostname", hostname, "--no-start")
							Eventually(session).Should(Say("routes:"))
							Eventually(session).Should(Say("(?i)\\+\\s+%s", route))
							Eventually(session).Should(Exit(0))
						})

						session := helpers.CF("app", appName)
						Eventually(session).Should(Say("name:\\s+%s", appName))
						Eventually(session).Should(Say("routes:\\s+%s", route))
						Eventually(session).Should(Exit(0))

						By("pushing a second time, without specifying a route")
						helpers.WithHelloWorldApp(func(dir string) {
							session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--no-start")
							Eventually(session).Should(Say("routes:"))
							Eventually(session).Should(Say("(?im)\\s+%s$", route))
							Eventually(session).Should(Exit(0))
						})
					})
				})

				Context("when at least one other route exists", func() {
					var (
						hostname2 string
						route2    string
					)

					BeforeEach(func() {
						helpers.WithHelloWorldApp(func(dir string) {
							session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--hostname", hostname, "--no-start")
							Eventually(session).Should(Exit(0))
						})

						hostname2 = strings.ToLower(helpers.RandomName())
						route2 = fmt.Sprintf("%s.%s", hostname2, helpers.DefaultSharedDomain())
					})

					It("creates and maps a route with a different hostname", func() {
						helpers.WithHelloWorldApp(func(dir string) {
							session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--hostname", hostname2, "--no-start")
							Eventually(session).Should(Say("routes:"))
							Eventually(session).Should(Say("(?i)\\+\\s+%s", route2))
							Eventually(session).Should(Exit(0))
						})

						session := helpers.CF("app", appName)
						Eventually(session).Should(Say("name:\\s+%s", appName))
						Eventually(session).Should(Say("routes:\\s+%s, %s", route, route2))
						Eventually(session).Should(Exit(0))
					})
				})
			})

			Context("the default route exists and is unmapped", func() {
				BeforeEach(func() {
					Eventually(helpers.CF("create-route", space, helpers.DefaultSharedDomain(), "-n", strings.ToLower(appName))).Should(Exit(0))
				})

				It("creates and maps the route with the provided hostname", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--hostname", hostname, "--no-start")
						Eventually(session).Should(Say("routes:"))
						Eventually(session).Should(Say("(?i)\\+\\s+%s", route))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Say("routes:\\s+%s", route))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the default route is mapped to the application", func() {
				BeforeEach(func() {
					helpers.WithHelloWorldApp(func(dir string) {
						Eventually(helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--no-start")).Should(Exit(0))
					})

					route = fmt.Sprintf("%s.%s", strings.ToLower(appName), helpers.DefaultSharedDomain())
				})

				It("does nothing", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "--hostname", hostname, "--no-start")
						Eventually(session).Should(Say("routes:"))
						Eventually(session).Should(Say("(?i)\\s+%s", route))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Say("routes:\\s+%s", route))
					Eventually(session).Should(Exit(0))
				})
			})
		})
	})

	Context("when using a tcp domain", func() {
		var (
			domain     helpers.Domain
			domainName string
		)

		BeforeEach(func() {
			domainName = helpers.DomainName()
			domain = helpers.NewDomain(organization, domainName)
			domain.CreateWithRouterGroup(helpers.FindOrCreateTCPRouterGroup(GinkgoParallelNode()))
		})

		AfterEach(func() {
			domain.DeleteShared()
		})

		It("returns an error", func() {
			session := helpers.CF(PushCommandName, appName, "--hostname", "I-dont-care", "-d", domainName, "--no-start")
			Eventually(session.Err).Should(Say("The route is invalid: a hostname cannot be used with a TCP domain."))
			Eventually(session).Should(Exit(1))
		})
	})
})
