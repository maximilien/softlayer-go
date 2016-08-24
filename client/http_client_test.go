package client_test

//TODO: need a fake Golang HTTPClient to implement this test.

//Public methods
// DoRawHttpRequestWithObjectMask
// DoRawHttpRequestWithObjectFilter
// DoRawHttpRequestWithObjectFilterAndObjectMask
// DoRawHttpRequest
// GenerateRequestBody
// HasErrors
// CheckForHttpResponseErrors

import (
	"bytes"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	slclient "github.com/maximilien/softlayer-go/client"
)

var _ = Describe("A HTTP Client", func() {
	var (
		//username, apiKey string
		server               *ghttp.Server
		client, wrongClient  *slclient.HttpClient
		err                  error
		endPoint             string
		slUsername, slAPIKey string
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		endPoint = strings.SplitAfter(server.URL(), "//")[1]
		slUsername = os.Getenv("SL_USERNAME")
		slAPIKey = os.Getenv("SL_API_KEY")
		client = slclient.NewHttpClient(slUsername, slAPIKey, endPoint, "templates", false)
		wrongClient = slclient.NewHttpClient(slUsername, slAPIKey, "10.0.0.0", "templates", false)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("#DoRawHttpRequest", func() {
		Context("when a successful request", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.VerifyRequest("GET", "/test"),
					ghttp.VerifyBasicAuth(slUsername, slAPIKey),
				)
			})

			It("make a request to access /test", func() {
				client.DoRawHttpRequest("test", "GET", bytes.NewBufferString("random text"))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(server.ReceivedRequests()).Should(HaveLen(1))
			})
		})

		Context("when getting i/o time error", func() {
			BeforeEach(func() {
				os.Setenv("SL_API_RETRY_COUNT", "2")
				os.Setenv("SL_API_WAIT_TIME", "1")
			})

			It("send a request to a wrong endPoint", func() {
				_, _, err := wrongClient.DoRawHttpRequest("test", "GET", bytes.NewBufferString("random text"))
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
