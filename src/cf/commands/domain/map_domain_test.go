package domain_test

import (
	"cf"
	. "cf/commands/domain"
	"cf/net"
	"errors"
	"github.com/stretchr/testify/assert"
	"testhelpers"
	"testing"
)

func TestMapDomainRequirements(t *testing.T) {
	ui := &testhelpers.FakeUI{}
	domainRepo := &testhelpers.FakeDomainRepository{}
	cmd := NewMapDomain(ui, domainRepo)

	ctxt := testhelpers.NewContext("map-domain", []string{"foo.com", "my-space"})
	reqFactory := &testhelpers.FakeReqFactory{LoginSuccess: true, TargetedOrgSuccess: true}

	testhelpers.RunCommand(cmd, ctxt, reqFactory)

	assert.True(t, testhelpers.CommandDidPassRequirements)

	reqFactory = &testhelpers.FakeReqFactory{LoginSuccess: true, TargetedOrgSuccess: false}
	testhelpers.RunCommand(cmd, ctxt, reqFactory)

	assert.False(t, testhelpers.CommandDidPassRequirements)

	reqFactory = &testhelpers.FakeReqFactory{LoginSuccess: false, TargetedOrgSuccess: true}
	testhelpers.RunCommand(cmd, ctxt, reqFactory)

	assert.False(t, testhelpers.CommandDidPassRequirements)

	ctxt = testhelpers.NewContext("map-domain", []string{})
	reqFactory = &testhelpers.FakeReqFactory{LoginSuccess: true, TargetedOrgSuccess: true}
	testhelpers.RunCommand(cmd, ctxt, reqFactory)

	assert.False(t, testhelpers.CommandDidPassRequirements)
}

func TestMapDomainSuccess(t *testing.T) {
	ctxt := testhelpers.NewContext("map-domain", []string{"foo.com", "my-space"})
	ui := &testhelpers.FakeUI{}
	domainRepo := &testhelpers.FakeDomainRepository{
		FindByNameInOrgDomain: cf.Domain{Name: "foo.com"},
	}

	reqFactory := &testhelpers.FakeReqFactory{
		LoginSuccess:       true,
		TargetedOrgSuccess: true,
		Organization:       cf.Organization{Name: "my-org", Guid: "my-org-guid"},
		Space:              cf.Space{Name: "my-space"},
	}

	cmd := NewMapDomain(ui, domainRepo)

	testhelpers.RunCommand(cmd, ctxt, reqFactory)
	assert.Equal(t, domainRepo.MapDomainDomain.Name, "foo.com")
	assert.Equal(t, domainRepo.MapDomainSpace.Name, "my-space")
	assert.Contains(t, ui.Outputs[0], "Mapping domain")
	assert.Contains(t, ui.Outputs[0], "foo.com")
	assert.Contains(t, ui.Outputs[0], "my-space")
	assert.Contains(t, ui.Outputs[1], "OK")
}

func TestMapDomainDomainNotFound(t *testing.T) {
	ctxt := testhelpers.NewContext("map-domain", []string{"foo.com", "my-space"})
	ui := &testhelpers.FakeUI{}
	domainRepo := &testhelpers.FakeDomainRepository{
		FindByNameInOrgApiStatus: net.NewNotFoundApiStatus(),
	}

	reqFactory := &testhelpers.FakeReqFactory{
		LoginSuccess:       true,
		TargetedOrgSuccess: true,
		Organization:       cf.Organization{Name: "my-org", Guid: "my-org-guid"},
		Space:              cf.Space{Name: "my-space"},
	}

	cmd := NewMapDomain(ui, domainRepo)

	testhelpers.RunCommand(cmd, ctxt, reqFactory)

	assert.Equal(t, len(ui.Outputs), 3)
	assert.Contains(t, ui.Outputs[0], "Mapping domain")
	assert.Contains(t, ui.Outputs[0], "foo.com")
	assert.Contains(t, ui.Outputs[0], "my-space")
	assert.Contains(t, ui.Outputs[1], "FAILED")
	assert.Contains(t, ui.Outputs[2], "foo.com")
	assert.Contains(t, ui.Outputs[2], "does not exist")
}

func TestMapDomainMappingFails(t *testing.T) {
	ctxt := testhelpers.NewContext("map-domain", []string{"foo.com", "my-space"})
	ui := &testhelpers.FakeUI{}
	domainRepo := &testhelpers.FakeDomainRepository{
		FindByNameInOrgDomain: cf.Domain{Name: "foo.com"},
		MapDomainApiStatus:    net.NewApiStatusWithError("Did not work %s", errors.New("bummer")),
	}

	reqFactory := &testhelpers.FakeReqFactory{
		LoginSuccess:       true,
		TargetedOrgSuccess: true,
		Organization:       cf.Organization{Name: "my-org", Guid: "my-org-guid"},
		Space:              cf.Space{Name: "my-space"},
	}

	cmd := NewMapDomain(ui, domainRepo)

	testhelpers.RunCommand(cmd, ctxt, reqFactory)

	assert.Equal(t, len(ui.Outputs), 3)
	assert.Contains(t, ui.Outputs[0], "Mapping domain")
	assert.Contains(t, ui.Outputs[0], "foo.com")
	assert.Contains(t, ui.Outputs[0], "my-space")
	assert.Contains(t, ui.Outputs[1], "FAILED")
	assert.Contains(t, ui.Outputs[2], "Did not work")
	assert.Contains(t, ui.Outputs[2], "bummer")
}
