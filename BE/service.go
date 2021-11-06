package profilesvc

import (
	"context"
	"errors"
)

type service interface {
	//membuat member
	Postmember(ctx context.Context, m member) error
	Getmember(ctx context.Context, id string) (member, error)
	
	//semua tentang artikel yang diakses member localhost:8080/member/id/artikel
	Postartikelmem(ctx context.Context, memberid string, a artikel) error
	Getartikelmem(ctx context.Context, memberid string) (artikel, error) //localhost:8080/member/id/artikel (list)
	Getartikel(ctx context.Context, memberid string, artikelid string) (artikel, error) //localhost:8080/member/id/artikel/id
	
	//semua tentang artikel yang diakses admin localhost:8080/admin/id/artikel
	Postartikeladm(ctx context.Context, adminid string, a artikel) error
	Getartikeladm(ctx context.Context, adminid string) (artikel, error) //localhost:8080/admin/id/artikel (list)
	Getartikeladmn(ctx context.Context, adminid string, artikelid string) (artikel, error) //localhost:8080/admin/id/artikel/id (artikel dipilih)

	//membuat admin
	Postadmin(ctx context.Context, aa admin) error
	Getadmin(ctx context.Context, id string) (admin, error)

	//semua tentang video yang diakses member localhost:8080/member/id/video
	Postvideomem(ctx context.Context, memberid string, v video) error
	Getvideomem(ctx context.Context, memberid string) (video, error)
	Getvideomember(ctx context.Context, memberid string, videoid string) (video, error)
	
	//semua tentang video yang diakses admin localhost:8080/admin/id/video
	Postvideoadm(ctx context.Context, adminid string, v video ) error
	Getvideoadm(ctx context.Context, adminid string) (video, error) //localhost:8080/admin/id/video (list)
	Getvideoadmn(ctx context.Context, adminid string, videoid string) (video, error) //localhost:8080/admin/id/video/id (video dipilih)

}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type inmemService struct {
	mtx sync.RWMutex
	m   map[string]Profile
}

func NewInmemService() Service {
	return &inmemService{
		m: map[string]Profile{},
	}
}

//buat member
func (s *inmemService) Postmember(ctx context.Context, m Profile) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[m.ID]; ok {
		return ErrAlreadyExists // POST = create, don't overwrite
	}
	s.m[m.ID] = m
	return nil
}

func (s *inmemService) Getmember(ctx context.Context, id string) (member, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	p, ok := s.m[id]
	if !ok {
		return member{}, ErrNotFound
	}
	return m, nil
}

//buat admin
func (s *inmemService) PostAdmin(ctx context.Context, aa admin) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[aa.ID]; ok {
		return ErrAlreadyExists // POST = create, don't overwrite
	}
	s.m[aa.ID] = aa
	return nil
}

func (s *inmemService) Getadmin(ctx context.Context, id string) (admin, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	aa, ok := s.m[id]
	if !ok {
		return admin{}, ErrNotFound
	}
	return aa, nil
}

//artikel member
func (s *inmemService) Getartikelmem(ctx context.Context, memberid string) (artikel, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	m, ok := s.m[memberid]
	if !ok {
		return artikel{}, ErrNotFound
	}
	return m.artikelmem, nil
}

func (s *inmemService) Getartikel(ctx context.Context, memberid string, artikelid string) (artikel, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	m, ok := s.m[memberid]
	if !ok {
		return artikel{}, ErrNotFound
	}
	for _, artikel := range m.artikelmem {
		if artikel.ID == artikelid {
			return artikel, nil
		}
	}
	return artikel{}, ErrNotFound
}

func (s *inmemService) Postartikelmem(ctx context.Context, memberid string, a artikel) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	m, ok := s.m[memberid]
	if !ok {
		return ErrNotFound
	}
	for _, artikel := range m.artikelmem {
		if artikel.ID == a.ID {
			return ErrAlreadyExists
		}
	}
	m.artikelmem = append(m.artikelmem, a)
	s.m[memberid] = m
	return nil
}

//artikel admin
func (s *inmemService) Getartikeladm(ctx context.Context, adminid string) (artikel, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	aa, ok := s.m[adminid]
	if !ok {
		return artikel{}, ErrNotFound
	}
	return m.artikeladm, nil
}

func (s *inmemService) Getartikeladmn(ctx context.Context, memberid string, artikelid string) (artikel, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	m, ok := s.m[memberid]
	if !ok {
		return artikel{}, ErrNotFound
	}
	for _, artikel := range m.artikeladm {
		if artikel.ID == artikelid {
			return artikel, nil
		}
	}
	return artikel{}, ErrNotFound
}

func (s *inmemService) Postartikelmem(ctx context.Context, memberid string, a artikel) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	aa, ok := s.m[adminid]
	if !ok {
		return ErrNotFound
	}
	for _, artikel := range m.artikeladm {
		if artikel.ID == a.ID {
			return ErrAlreadyExists
		}
	}
	aa.artikeladm = append(aa.artikeladm, a)
	s.m[adminid] = aa
	return nil
}

//video member
func (s *inmemService) Getvideomem(ctx context.Context, memberid string) (video, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	m, ok := s.m[memberid]
	if !ok {
		return video{}, ErrNotFound
	}
	return m.videomem, nil
}

func (s *inmemService) Getvideomember(ctx context.Context, memberid string, videoid string) (video, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	m, ok := s.m[memberid]
	if !ok {
		return video{}, ErrNotFound
	}
	for _, video := range m.videomem {
		if videoid == videoid {
			return video, nil
		}
	}
	return video{}, ErrNotFound
}

func (s *inmemService) Post(ideomemctx context.Context, memberid string v video) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	pm ok := s.m[pmemberid
	if !ok {
		return ErrNotFound
	}
	for _, avideo:= range pv.videomem{
		if avideoID == avID {
			return ErrAlreadyExists
		}
	}
	pmAvideomem= append(pm.videomem av
	s.m[pmemberid = pm
	return nil
}



//video admin
func (s *inmemService) Getvideoadm(ctx context.Context, adminid string) (video, error) {
    s.mtx.RLock()
    defer s.mtx.RUnlock()
    aa, ok := s.m[adminid]
    if !ok {
        return []video{}, ErrNotFound
    }
    return aa.videoadm nil
}

func (s *inmemService) Getvideoadmn(ctx context.Context, adminid string, videoid string) (video error) {
    s.mtx.RLock()
    defer s.mtx.RUnlock()
    aa ok := s.m[adminid]
    if !ok {
        return videoadm{}, ErrNotFound
    }
    for _, address := range aa.videoadm {
        if videoid == videoid {
            return video, nil
        }
    }
    return videoadm{}, ErrNotFound
}

func (s *inmemService) Postvideoadm(ctx context.Context, adminid string, v video) error {
    s.mtx.Lock()
    defer s.mtx.Unlock()
    aa, ok := s.adminid]
    if !ok {
        return ErrNotFound
    }
    for _,video := range aa.videoadm {
        if video.ID == v.ID {
            return ErrAlreadyExists
        }
    }
    aa.videoadm = append(aa.videoadm, a)
    s.m[adminid] = aa
    return nil
}