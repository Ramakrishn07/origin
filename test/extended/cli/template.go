package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/util/wait"
	admissionapi "k8s.io/pod-security-admission/api"

	exutil "github.com/openshift/origin/test/extended/util"
)

var _ = g.Describe("[sig-cli] templates", func() {
	defer g.GinkgoRecover()

	var (
		oc                           = exutil.NewCLIWithPodSecurityLevel("oc-templates", admissionapi.LevelBaseline)
		testDataPath                 = exutil.FixturePath("testdata", "cmd", "test", "cmd", "testdata")
		appTemplatePath              = filepath.Join(testDataPath, "application-template-dockerbuild.json")
		appTemplateStiPath           = filepath.Join(testDataPath, "application-template-stibuild.json")
		guestbookTemplatePath        = filepath.Join(testDataPath, "templates", "guestbook.json")
		guestbookTemplateEnvPath     = filepath.Join(testDataPath, "templates", "guestbook.env")
		templateRequiredParamPath    = filepath.Join(testDataPath, "templates", "template_required_params.yaml")
		templateRequiredParamEnvPath = filepath.Join(testDataPath, "templates", "template_required_params.env")
		templateTypePrecisionPath    = filepath.Join(testDataPath, "templates", "template-type-precision.json")
		basicUsersBindingPath        = filepath.Join(testDataPath, "templates", "basic-users-binding.json")
		multilinePath                = filepath.Join(testDataPath, "templates", "multiline.txt")
	)

	g.It("process [apigroup:apps.openshift.io][apigroup:template.openshift.io][Skipped:Disconnected]", func() {
		err := oc.Run("get").Args("templates").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("create").Args("-f", appTemplatePath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("get").Args("templates").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("get").Args("templates", "ruby-helloworld-sample").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		outputYamlFile, err := oc.WithoutNamespace().Run("get").Args("template", "ruby-helloworld-sample", "-o", "json").OutputToFile("template.json")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.WithoutNamespace().Run("process").Args("-f", outputYamlFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("process").Args("ruby-helloworld-sample").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err := oc.Run("process").Args("ruby-helloworld-sample", "-o", "template", "--template", "{{.kind}}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.Equal("List"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "go-template", "--template", "{{.kind}}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.Equal("List"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "go-template={{.kind}}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.Equal("List"))
		err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "go-template-file=/dev/null").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "jsonpath", "--template", "{.kind}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.Equal("List"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "jsonpath={.kind}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.Equal("List"))
		err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "jsonpath-file=/dev/null").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "describe").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("ruby-27-centos7"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "json").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("ruby-27-centos7"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "yaml").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("ruby-27-centos7"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-o", "name").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("ruby-27-centos7"))
		out, err = oc.Run("describe").Args("templates", "ruby-helloworld-sample").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("BuildConfig"))
		o.Expect(out).To(o.ContainSubstring("ruby-sample-build"))
		err = oc.Run("delete").Args("-f", appTemplatePath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		outputTemplateFile, err := oc.Run("process").Args("-f", guestbookTemplatePath, "-l", "app=guestbook").OutputToFile("guestbooktemplate.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("create").Args("-f", outputTemplateFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("status").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("frontend-service"))
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, "--local", "-l", "app=guestbook", "-o", "yaml").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("app: guestbook"))

		localOut, err := oc.Run("process").Args("-f", guestbookTemplatePath, "--local", "-l", "app=guestbook", "-o", "yaml", "ADMIN_USERNAME=au", "ADMIN_PASSWORD=ap", "REDIS_PASSWORD=rp").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		remoteOut, err := oc.Run("process").Args("-f", guestbookTemplatePath, "-l", "app=guestbook", "-o", "yaml", "ADMIN_USERNAME=au", "ADMIN_PASSWORD=ap", "REDIS_PASSWORD=rp").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(localOut).To(o.Equal(remoteOut))

		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, "--local", "-l", "app=guestbook", "-o", "yaml", "--server", "0.0.0.0:1").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("app: guestbook"))
		err = oc.Run("delete").Args("-f", outputTemplateFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		g.By("individually specified parameter values are honored")
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, "-p", "ADMIN_USERNAME=myuser", "-p", "ADMIN_PASSWORD=mypassword").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("myuser"))
		o.Expect(out).To(o.ContainSubstring("mypassword"))
		out, err = oc.Run("process").Args("ADMIN_USERNAME=myuser", "ADMIN_PASSWORD=mypassword", "-f", guestbookTemplatePath).Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("myuser"))
		o.Expect(out).To(o.ContainSubstring("mypassword"))
		err = oc.Run("create").Args("-f", appTemplateStiPath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "MYSQL_USER=myself", "MYSQL_PASSWORD=my,1%pa=s").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("myself"))
		out, err = oc.Run("process").Args("MYSQL_USER=myself", "MYSQL_PASSWORD=my,1%pa=s", "ruby-helloworld-sample").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("my,1%pa=s"))
		out, err = oc.Run("process").Args("ruby-helloworld-sample", "-p", "MYSQL_USER=myself", "-p", "MYSQL_PASSWORD=my,1%pa=s").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("myself"))
		out, err = oc.Run("process").Args("-p", "MYSQL_USER=myself", "-p", "MYSQL_PASSWORD=my,1%pa=s", "ruby-helloworld-sample").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("my,1%pa=s"))
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, fmt.Sprintf("--param-file=%s", guestbookTemplateEnvPath)).Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("root"))
		o.Expect(out).To(o.ContainSubstring("adminpass"))
		o.Expect(out).To(o.ContainSubstring("redispass"))
		templateEnvs, err := os.ReadFile(guestbookTemplateEnvPath)
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, "--param-file=-").InputString(string(templateEnvs)).Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("root"))
		o.Expect(out).To(o.ContainSubstring("adminpass"))
		o.Expect(out).To(o.ContainSubstring("redispass"))
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, fmt.Sprintf("--param-file=%s", guestbookTemplateEnvPath), "-p", "ADMIN_USERNAME=myuser").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("ignoring value from file"))
		o.Expect(out).To(o.ContainSubstring("myuser"))
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, fmt.Sprintf("--param-file=%s", guestbookTemplateEnvPath), "-p", "ADMIN_PASSWORD=mypassword").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("mypassword"))
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, fmt.Sprintf("--param-file=%s", guestbookTemplateEnvPath), "-p", "REDIS_PASSWORD=rrr").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("rrr"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, fmt.Sprintf("--param-file=%s", templateRequiredParamEnvPath), "-o", "yaml").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("first"))
		o.Expect(out).To(o.ContainSubstring("second"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param-file=does/not/exist").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("no such file or directory"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, fmt.Sprintf("--param-file=%s", testDataPath)).Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("is a directory"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param-file=/dev/null").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("parameter required_param is required and must be specified"))
		err = oc.Run("process").Args("-f", guestbookTemplatePath, "--param-file=/dev/null", fmt.Sprintf("--param-file=%s", guestbookTemplateEnvPath)).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("-f", guestbookTemplatePath, "-p", "ABSENT_PARAMETER=absent").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("unknown parameter name"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param-file=-").InputString("fo%(o=bar").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("invalid parameter assignment"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param-file=-").InputString("S P A C E S=test").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("invalid parameter assignment"))
		err = oc.Run("process").Args("-f", guestbookTemplatePath, "-p", "ABSENT_PARAMETER=absent", "--ignore-unknown-parameters").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("delete").Args("-f", appTemplateStiPath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		g.By("Ensure large integers survive unstructured JSON creation")
		err = oc.Run("create").Args("-f", templateTypePrecisionPath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("template-type-precision").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("1000030003"))
		o.Expect(out).To(o.ContainSubstring("2147483647"))
		o.Expect(out).To(o.ContainSubstring("9223372036854775807"))
		outPrecisionFile, err := oc.Run("process").Args("template-type-precision").OutputToFile("template-type-precision.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("create").Args("-f", outPrecisionFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("get").Args("pod/template-type-precision", "-o", "json").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("1000030003"))
		o.Expect(out).To(o.ContainSubstring("2147483647"))
		o.Expect(out).To(o.ContainSubstring("9223372036854775807"))
		patch := `{"metadata":{"annotations":{"comment":"patch comment"}}}`
		err = oc.Run("patch").Args("pod", "template-type-precision", "-p", patch).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("get").Args("pod/template-type-precision", "-o", "json").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("9223372036854775807"))
		o.Expect(out).To(o.ContainSubstring("patch comment"))
		err = oc.Run("delete").Args("template/template-type-precision").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("delete").Args("pod/template-type-precision").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		g.By("validates oc process")
		out, err = oc.Run("process").Args("name1", "name2").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("template name must be specified only once"))
		out, err = oc.Run("process").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("Must pass a filename or name of stored template"))

		g.By("can't ask for parameters and try process the template")
		out, err = oc.Run("process").Args("template-name", "--parameters", "--param=someval").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --param"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "-p", "someval").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --param"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "--labels=someval").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --labels"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "-l", "someval").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --labels"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "--output=yaml").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --output"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "-o", "yaml").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --output"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "--raw").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --raw"))
		out, err = oc.Run("process").Args("template-name", "--parameters", "--template=someval").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("--parameters flag does not process the template, can't be used with --template"))

		g.By("providing a value more than once should fail")
		out, err = oc.Run("process").Args("template-name", "key=value", "key=value").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("provided more than once: key"))
		out, err = oc.Run("process").Args("template-name", "--param=key=value", "--param=key=value").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("provided more than once: key"))
		out, err = oc.Run("process").Args("template-name", "key=value", "--param=key=value").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("provided more than once: key"))
		out, err = oc.Run("process").Args("template-name", "key=value", "other=foo", "--param=key=value", "--param=other=baz").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("provided more than once: key, other"))
		out, err = oc.Run("process").Args("-f", basicUsersBindingPath).Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("not a valid Template but"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath).Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("parameter required_param is required and must be specified"))
		err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=someval").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		testUser := oc.Namespace() + "-someval-user"
		requiredParamFile, err := oc.Run("process").Args("-f", templateRequiredParamPath, "-p", fmt.Sprintf("required_param=%s", testUser)).OutputToFile("required_param.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.AsAdmin().Run("create").Args("-f", requiredParamFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		defer func() {
			oc.AsAdmin().Run("delete").Args("user", testUser).Execute()
		}()
		requiredParamFile, err = oc.Run("process").Args("-f", templateRequiredParamPath, fmt.Sprintf("required_param=%s=moreval", testUser)).OutputToFile("required_param.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.AsAdmin().Run("create").Args("-f", requiredParamFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		defer func() {
			oc.AsAdmin().Run("delete").Args("user", testUser+"=moreval").Execute()
		}()
		requiredParamFile, err = oc.Run("process").Args("-f", templateRequiredParamPath, "-p", fmt.Sprintf("required_param=%s=moreval2", testUser)).OutputToFile("required_param.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.AsAdmin().Run("create").Args("-f", requiredParamFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		defer func() {
			oc.AsAdmin().Run("delete").Args("user", testUser+"=moreval2").Execute()
		}()
		requiredParamFile, err = oc.Run("process").Args("-f", templateRequiredParamPath, "-p", fmt.Sprintf("required_param=%s=moreval3", testUser)).OutputToFile("required_param.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.AsAdmin().Run("create").Args("-f", requiredParamFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		defer func() {
			oc.AsAdmin().Run("delete").Args("user", testUser+"=moreval3").Execute()
		}()
		out, err = oc.AsAdmin().Run("get").Args("user", testUser, "-o", "jsonpath={.metadata.name}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring(testUser))
		out, err = oc.AsAdmin().Run("get").Args("user", testUser+"=moreval", "-o", "jsonpath={.metadata.name}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring(testUser + "=moreval"))
		out, err = oc.AsAdmin().Run("get").Args("user", testUser+"=moreval2", "-o", "jsonpath={.metadata.name}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring(testUser + "=moreval2"))
		out, err = oc.AsAdmin().Run("get").Args("user", testUser+"=moreval3", "-o", "jsonpath={.metadata.name}").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring(testUser + "=moreval3"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=someval", "--param=other_param=otherval").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("unknown parameter name \"other_param\""))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=someval", "--param=optional_param").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("invalid parameter assignment in"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=someval", "--labels======").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("error parsing labels"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=optional_param=a,required_param=b").Output()
		o.Expect(err).To(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("parameter required_param is required and must be specified"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=a,b=c,d").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("no longer accepts comma-separated list"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=a_b_c_d").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).NotTo(o.ContainSubstring("no longer accepts comma-separated list"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "--param=required_param=a,b,c,d").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).NotTo(o.ContainSubstring("no longer accepts comma-separated list"))
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, "required_param=a,b=c,d").Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).NotTo(o.ContainSubstring("no longer accepts comma-separated list"))
		multiLine, err := os.ReadFile(multilinePath)
		o.Expect(err).NotTo(o.HaveOccurred())
		out, err = oc.Run("process").Args("-f", templateRequiredParamPath, fmt.Sprintf("--param=required_param=%s", string(multiLine))).Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring("also,with=commas"))
	})

	g.It("different namespaces [apigroup:user.openshift.io][apigroup:project.openshift.io][apigroup:template.openshift.io][apigroup:authorization.openshift.io][Skipped:Disconnected]", func() {
		bob := oc.CreateUser("bob-")

		err := oc.Run("create").Args("-f", appTemplatePath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.Run("policy").Args("add-role-to-user", "admin", bob.Name).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		oc.ChangeUser(bob.Name)

		testProject2 := oc.Namespace() + "-project2"
		out, err := oc.Run("new-project").Args(testProject2).Output()
		o.Expect(err).NotTo(o.HaveOccurred())
		o.Expect(out).To(o.ContainSubstring(fmt.Sprintf("Now using project \"%s\" on server ", testProject2)))
		defer func() {
			err = oc.WithoutNamespace().Run("delete", "project").Args(testProject2).Execute()
			o.Expect(err).NotTo(o.HaveOccurred())
		}()

		err = wait.PollImmediate(500*time.Millisecond, time.Minute, func() (bool, error) {
			return oc.WithoutNamespace().Run("get").Args("templates").Execute() == nil, nil
		})
		o.Expect(err).NotTo(o.HaveOccurred())

		err = oc.WithoutNamespace().Run("create").Args("-f", appTemplatePath).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		err = oc.WithoutNamespace().Run("process").Args("template/ruby-helloworld-sample").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.WithoutNamespace().Run("process").Args("templates/ruby-helloworld-sample").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.WithoutNamespace().Run("process").Args(oc.Namespace() + "//ruby-helloworld-sample").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
		err = oc.WithoutNamespace().Run("process").Args(oc.Namespace() + "/template/ruby-helloworld-sample").Execute()
		o.Expect(err).NotTo(o.HaveOccurred())

		outputYamlFile, err := oc.WithoutNamespace().Run("get").Args("template", "ruby-helloworld-sample", "-o", "yaml").OutputToFile("template.yaml")
		o.Expect(err).NotTo(o.HaveOccurred())

		err = oc.WithoutNamespace().Run("process").Args("-f", outputYamlFile).Execute()
		o.Expect(err).NotTo(o.HaveOccurred())
	})
})
