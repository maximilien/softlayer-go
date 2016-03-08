package client_test

// import (
// 	"os"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"

// 	slclient "github.com/maximilien/softlayer-go/client"
// 	softlayer "github.com/maximilien/softlayer-go/softlayer"
// )

// var _ = Describe("HttpClient", func() {
// 	var (
// 		username, apiKey string
// 		client           softlayer.HttpClient
// 	)

// 	BeforeEach(func() {
// 		os.Setenv("SL_GO_NON_VERBOSE", "TRUE")

// 		client = slclient.NewHttpClient(username, apiKey)
// 	})

// 	// Context("#NewSoftLayerClient", func() {
// 	// 	It("creates a new client with username and apiKey", func() {
// 	// 		client = slclient.NewSoftLayerClient(username, apiKey)
// 	// 		Expect(client).ToNot(BeNil())
// 	// 	})
// 	// })

// 	// Context("#NewSoftLayerClient_HTTPClient", func() {
// 	// 	It("creates a new client which should have an initialized default HTTP client", func() {
// 	// 		client = slclient.NewSoftLayerClient(username, apiKey)
// 	// 		if c, ok := client.(*slclient.SoftLayerClient); ok {
// 	// 			Expect(c.HTTPClient).ToNot(BeNil())
// 	// 		}
// 	// 	})

// 	// 	It("creates a new client with a custom HTTP client", func() {
// 	// 		client = slclient.NewSoftLayerClient(username, apiKey)
// 	// 		c, _ := client.(*slclient.SoftLayerClient)

// 	// 		// Assign a malformed dialer to test if the HTTP client really works
// 	// 		var errDialFailed = errors.New("dial failed")
// 	// 		c.HTTPClient = &http.Client{
// 	// 			Transport: &http.Transport{
// 	// 				Dial: func(network, addr string) (net.Conn, error) {
// 	// 					return nil, errDialFailed
// 	// 				},
// 	// 			},
// 	// 		}

// 	// 		_, err := c.DoRawHttpRequest("/foo", "application/text", bytes.NewBufferString("random text"))
// 	// 		Expect(err).To(Equal(errDialFailed))

// 	// 	})
// 	// })

// 	Context("#IsHttpErrorCode", func() {
// 		Expect(true).To(BeTrue())
// 		//TODO: add tests
// 	})

// 	Context("#DoRawHttpRequestWithObjectMask", func() {
// 		//TODO: add tests
// 	})

// 	Context("#DoRawHttpRequestWithObjectFilter", func() {
// 		//TODO: add tests
// 	})

// 	Context("#DoRawHttpRequest", func() {
// 		//TODO: add tests
// 	})

// 	Context("#DoRawHttpRequestWithObjectFilterAndObjectMask", func() {
// 		//TODO: add tests
// 	})
// })
