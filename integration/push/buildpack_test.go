package push

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"

	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("push with different buildpack values", func() {
	var (
		appName string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()
	})

	Context("when the buildpack flag is provided", func() {
		Context("when only one buildpack is provided", func() {
			It("sets that buildpack correctly for the pushed app", func() {
				helpers.WithProcfileApp(func(dir string) {
					tempfile := filepath.Join(dir, "index.html")
					err := ioutil.WriteFile(tempfile, []byte(fmt.Sprintf("hello world %d", rand.Int())), 0666)
					Expect(err).ToNot(HaveOccurred())

					By("pushing a ruby app with a static buildpack sets buildpack to static")
					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
						PushCommandName, appName,
						"-b", "staticfile_buildpack",
					)
					Eventually(session).Should(Say("Staticfile Buildpack"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("app", appName)
					Eventually(session).Should(Say("buildpack:\\s+staticfile_buildpack"))
					Eventually(session).Should(Exit(0))

					By("pushing a ruby app with a null buildpack sets buildpack to auto-detected (ruby)")
					session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
						PushCommandName, appName,
						"-b", "null",
					)
					Eventually(session).Should(Say(`\-\s+buildpack:\s+staticfile_buildpack`))
					Consistently(session).ShouldNot(Say(`\+\s+buildpack:`))
					Eventually(session).Should(Say("Ruby Buildpack"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("app", appName)
					Eventually(session).Should(Say(`buildpack:\s+ruby`))
					Eventually(session).Should(Exit(0))

					By("pushing a ruby app with a static buildpack sets buildpack to static")
					session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
						PushCommandName, appName,
						"-b", "staticfile_buildpack",
					)
					Eventually(session).Should(Say("Staticfile Buildpack"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("app", appName)
					Eventually(session).Should(Say("buildpack:\\s+staticfile_buildpack"))
					Eventually(session).Should(Exit(0))

					By("pushing a ruby app with a default buildpack sets buildpack to auto-detected (ruby)")
					session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
						PushCommandName, appName,
						"-b", "default",
					)
					Eventually(session).Should(Say(`\-\s+buildpack:\s+staticfile_buildpack`))
					Consistently(session).ShouldNot(Say(`\+\s+buildpack:`))
					Eventually(session).Should(Say("Ruby Buildpack"))
					Eventually(session).Should(Exit(0))

					session = helpers.CF("app", appName)
					Eventually(session).Should(Say(`buildpack:\s+ruby`))
					Eventually(session).Should(Exit(0))

				})
			})
		})

		Context("when multiple instances of buildpack are provided", func() {
			Context("when the app does NOT have existing buildpack configurations", func() {
				It("pushes the app successfully with multiple buildpacks", func() {
					helpers.WithProcfileApp(func(dir string) {
						tempfile := filepath.Join(dir, "index.html")
						err := ioutil.WriteFile(tempfile, []byte(fmt.Sprintf("hello world %d", rand.Int())), 0666)
						Expect(err).ToNot(HaveOccurred())

						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
							PushCommandName, appName,
							"-b", "staticfile_buildpack", "-b", "ruby_buildpack", "--no-start",
						)
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("curl", fmt.Sprintf("v3/apps/%s", helpers.AppGUID(appName)))

					Eventually(session).Should(Say(`\s+"buildpacks":\s+`))
					Eventually(session).Should(Say(`\s+"staticfile_buildpack"`))
					Eventually(session).Should(Say(`\s+"ruby_buildpack"`))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the app has existing buildpacks", func() {
				It("pushes the app successfully and overrides the existing buildpacks", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
							"applications": []map[string]interface{}{
								{
									"name": appName,
									"buildpacks": []string{
										"ruby_buildpack",
										"staticfile_buildpack",
									},
								},
							},
						})
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
							PushCommandName, appName,
							"-b", "php_buildpack", "-b", "go_buildpack", "--no-start",
						)
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("curl", fmt.Sprintf("v3/apps/%s", helpers.AppGUID(appName)))

					Eventually(session).Should(Say(`\s+"buildpacks":\s+`))
					Eventually(session).Should(Say(`php_buildpack`))
					Eventually(session).Should(Say(`go_buildpack`))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the app has existing `buildpack`", func() {
				It("pushes the app successfully and overrides the existing buildpacks", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
							"applications": []map[string]interface{}{
								{
									"name": appName,
									"buildpacks": []string{
										"staticfile_buildpack",
									},
								},
							},
						})
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
							PushCommandName, appName,
							"-b", "php_buildpack", "-b", "go_buildpack", "--no-start",
						)
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("curl", fmt.Sprintf("v3/apps/%s", helpers.AppGUID(appName)))

					Eventually(session).Should(Say(`\s+"buildpacks":\s+`))
					Eventually(session).Should(Say(`php_buildpack`))
					Eventually(session).Should(Say(`go_buildpack`))
					Eventually(session).Should(Exit(0))
				})
			})
		})
	})

	Context("when buildpack is provided via manifest", func() {
		It("sets buildpack and returns a warning", func() {
			helpers.WithHelloWorldApp(func(dir string) {
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name":      appName,
							"buildpack": "staticfile_buildpack",
						},
					},
				})
				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName, "no-start")
				Eventually(session).Should(Say(`\s+buildpack:\s+staticfile_buildpack`))
				Eventually(session.Err).Should(Say(`Deprecation warning: Use of buildpack`))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when buildpacks (plural) is provided via manifest", func() {
		Context("when mutiple buildpacks are specified", func() {
			It("sets all buildpacks correctly for the pushed app", func() {
				helpers.WithHelloWorldApp(func(dir string) {
					helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
						"applications": []map[string]interface{}{
							{
								"name": appName,
								"buildpacks": []string{
									"https://github.com/cloudfoundry/ruby-buildpack",
									"https://github.com/cloudfoundry/staticfile-buildpack",
								},
							},
						},
					})
					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
					Eventually(session).Should(Exit(0))
				})

				session := helpers.CF("curl", fmt.Sprintf("v3/apps/%s", helpers.AppGUID(appName)))

				Eventually(session).Should(Say(`https://github.com/cloudfoundry/ruby-buildpack"`))
				Eventually(session).Should(Say(`https://github.com/cloudfoundry/staticfile-buildpack"`))
				Eventually(session).Should(Exit(0))
			})
		})

		Context("when only one buildpack is specified", func() {
			It("sets only one buildpack for the pushed app", func() {
				helpers.WithHelloWorldApp(func(dir string) {
					helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
						"applications": []map[string]interface{}{
							{
								"name": appName,
								"buildpacks": []string{
									"https://github.com/cloudfoundry/staticfile-buildpack",
								},
							},
						},
					})
					session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
					Eventually(session).Should(Exit(0))
				})

				session := helpers.CF("curl", fmt.Sprintf("v3/apps/%s", helpers.AppGUID(appName)))

				// TODO: fix during app command rework to actually test that the second buildpack does not exist
				Eventually(session).Should(Say(`https://github.com/cloudfoundry/staticfile-buildpack"`))
				Eventually(session).Should(Exit(0))
			})
		})
	})
})
