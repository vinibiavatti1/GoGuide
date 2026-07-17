// Proxy
// The Proxy pattern is a structural design pattern that provides an object representing another object.
// It acts as a surrogate or placeholder for another object to control access to it.

package structural

// Protocol
// We will define this interface to represent a common protocol.
// This will be implemented by the main service, and by the proxy.
type DataAccess interface {
	Query(query string) string
}

// Service
// This is a concrete implementation of the DataAccess interface.
// It represents the main service that will be used to access the data.
type DataService struct{}

// Service Implementation
// This method will be used to query the data.
func (d *DataService) Query(query string) string {
	return "data"
}

// Proxy
// The struct below is a Proxy.
// It has a reference to the main service, and implements the same interface.
// It will be used to control access to the main service.
// In this case, it will cache the results of the queries, and will call the main service only if the result
// is not in the cache.
type CachedDataService struct {
	Cache   map[string]string
	Service DataAccess
}

// Proxy Implementation
// This is the same method from the main service, however, it will check if the result is in the cache.
// If it is, it will return the cached result.
func (d *CachedDataService) Query(query string) string {
	res, ok := d.Cache[query]
	if ok {
		return res
	}
	res = d.Service.Query(query)
	d.Cache[query] = res
	return res
}

// Test Proxy
// The test function will create a main service and a proxy.
// It will call the main service and the proxy.
// Note that when using the main service directly, the result is always computed.
// When using the proxy, the result is cached after the first call.
func TestProxy() {
	ds := &DataService{}
	ds.Query("abc") // Computed
	ds.Query("abc") // Computed

	cds := &CachedDataService{
		Cache:   map[string]string{},
		Service: ds,
	}
	cds.Query("abc") // Computed
	cds.Query("abc") // From Cache
}
