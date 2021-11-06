package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	PostProfileEndpoint endpoint.Endpoint
	GetProfileEndpoint  endpoint.Endpoint

	GetAddressesEndpoint endpoint.Endpoint
	GetAddressEndpoint   endpoint.Endpoint
	PostAddressEndpoint  endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostProfileEndpoint: MakePostProfileEndpoint(s),
		GetProfileEndpoint:  MakeGetProfileEndpoint(s),

		GetAddressesEndpoint: MakeGetAddressesEndpoint(s),
		GetAddressEndpoint:   MakeGetAddressEndpoint(s),
		PostAddressEndpoint:  MakePostAddressEndpoint(s),
	}
}

// Postmember(ctx context.Context, m member) error
// 	Getmember(ctx context.Context, id string) (member, error)

// 	//semua tentang artikel yang diakses member localhost:8080/member/id/artikel
// 	Postartikelmem(ctx context.Context, memberid string, a artikel) error
// 	Getartikelmem(ctx context.Context, memberid string) (artikel, error) //localhost:8080/member/id/artikel (list)
// 	Getartikel(ctx context.Context, memberid string, artikelid string) (artikel, error) //localhost:8080/member/id/artikel/id

// 	//semua tentang artikel yang diakses admin localhost:8080/admin/id/artikel
// 	Postartikeladm(ctx context.Context, adminid string, a artikel) error
// 	Getartikeladm(ctx context.Context, adminid string) (artikel, error) //localhost:8080/admin/id/artikel (list)
// 	Getartikeladmn(ctx context.Context, adminid string, artikelid string) (artikel, error) //localhost:8080/admin/id/artikel/id (artikel dipilih)

// 	//membuat admin
// 	Postadmin(ctx context.Context, aa admin) error
// 	Getadmin(ctx context.Context, id string) (admin, error)

// 	//semua tentang video yang diakses member localhost:8080/member/id/video
// 	Postvideomem(ctx context.Context, memberid string, v video) error
// 	Getvideomem(ctx context.Context, memberid string) (video, error)
// 	Getvideomember(ctx context.Context, memberid string, videoid string) (video, error)

// 	//semua tentang video yang diakses admin localhost:8080/admin/id/video
// 	Postvideoadm(ctx context.Context, adminid string, v video ) error
// 	Getvideoadm(ctx context.Context, adminid string) (video, error) //localhost:8080/admin/id/video (list)
// 	Getvideoadmn(ctx context.Context, adminid string, videoid string) (video, error) //localhost:8080/admin/id/video/id (video dipilih)

// PostProfile implements Service. Primarily useful in a client.
func (e Endpoints) PostProfile(ctx context.Context, p Profile) error {
	request := postProfileRequest{Profile: p}
	response, err := e.PostProfileEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postProfileResponse)
	return resp.Err
}

// GetProfile implements Service. Primarily useful in a client.
func (e Endpoints) GetProfile(ctx context.Context, id string) (Profile, error) {
	request := getProfileRequest{ID: id}
	response, err := e.GetProfileEndpoint(ctx, request)
	if err != nil {
		return Profile{}, err
	}
	resp := response.(getProfileResponse)
	return resp.Profile, resp.Err
}

// GetAddresses implements Service. Primarily useful in a client.
func (e Endpoints) GetAddresses(ctx context.Context, profileID string) ([]Address, error) {
	request := getAddressesRequest{ProfileID: profileID}
	response, err := e.GetAddressesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(getAddressesResponse)
	return resp.Addresses, resp.Err
}

// GetAddress implements Service. Primarily useful in a client.
func (e Endpoints) GetAddress(ctx context.Context, profileID string, addressID string) (Address, error) {
	request := getAddressRequest{ProfileID: profileID, AddressID: addressID}
	response, err := e.GetAddressEndpoint(ctx, request)
	if err != nil {
		return Address{}, err
	}
	resp := response.(getAddressResponse)
	return resp.Address, resp.Err
}

// PostAddress implements Service. Primarily useful in a client.
func (e Endpoints) PostAddress(ctx context.Context, profileID string, a Address) error {
	request := postAddressRequest{ProfileID: profileID, Address: a}
	response, err := e.PostAddressEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postAddressResponse)
	return resp.Err
}

// MakeGetAddressesEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.

func MakePostProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postProfileRequest)
		e := s.PostProfile(ctx, req.Profile)
		return postProfileResponse{Err: e}, nil
	}
}

// MakeGetProfileEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getProfileRequest)
		p, e := s.GetProfile(ctx, req.ID)
		return getProfileResponse{Profile: p, Err: e}, nil
	}
}

func MakeGetAddressesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getAddressesRequest)
		a, e := s.GetAddresses(ctx, req.ProfileID)
		return getAddressesResponse{Addresses: a, Err: e}, nil
	}
}

// MakeGetAddressEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getAddressRequest)
		a, e := s.GetAddress(ctx, req.ProfileID, req.AddressID)
		return getAddressResponse{Address: a, Err: e}, nil
	}
}

// MakePostAddressEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePostAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postAddressRequest)
		e := s.PostAddress(ctx, req.ProfileID, req.Address)
		return postAddressResponse{Err: e}, nil
	}
}

type postProfileRequest struct {
	Profile Profile
}

type postProfileResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postProfileResponse) error() error { return r.Err }

type getProfileRequest struct {
	ID string
}

type getProfileResponse struct {
	Profile Profile `json:"profile,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func (r getProfileResponse) error() error { return r.Err }

type getAddressesRequest struct {
	ProfileID string
}

type getAddressesResponse struct {
	Addresses []Address `json:"addresses,omitempty"`
	Err       error     `json:"err,omitempty"`
}

func (r getAddressesResponse) error() error { return r.Err }

type getAddressRequest struct {
	ProfileID string
	AddressID string
}

type getAddressResponse struct {
	Address Address `json:"address,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func (r getAddressResponse) error() error { return r.Err }

type postAddressRequest struct {
	ProfileID string
	Address   Address
}

type postAddressResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postAddressResponse) error() error { return r.Err }
