package pwapi

import (
	"context"
	"testing"

	"moul.io/godev"
	"pathwar.land/go/internal/testutil"
	"pathwar.land/go/pkg/pwsso"
)

func TestSvc_UserGetSession(t *testing.T) {
	svc, cleanup := TestingService(t, ServiceOpts{Logger: testutil.Logger(t)})
	defer cleanup()
	ctx := testingSetContextToken(context.Background(), t)

	session, err := svc.UserGetSession(ctx, nil)
	checkErr(t, "", err)

	//fmt.Println(godev.PrettyJSON(session))
	var tests = []struct {
		name     string
		actual   interface{}
		expected string
	}{
		{".User.Username", session.User.Username, `"moul"`},
		{"len(.Season)", len(session.Seasons), "2"},
		{".Claims", session.Claims, godev.JSON(pwsso.TestingClaims(t))},
		{".IsNewUser", session.IsNewUser, `true`},
		{".User.ActiveTeamMember.Team.Season.Name", session.User.ActiveTeamMember.Team.Season.Name, `"Solo Mode"`},
		{".User.ActiveTeamMember.Team.Organization.Name", session.User.ActiveTeamMember.Team.Organization.Name, `"moul"`},
		{".User.ActiveTeamMember.Team.Organization.SoloSeason", session.User.ActiveTeamMember.Team.Organization.SoloSeason, `true`},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := godev.JSON(test.actual)
			if test.expected != actual {
				t.Fatalf("Expected: %q, got %q instead.", test.expected, actual)
			}
		})
	}
}