package routes

import (
	"html/template"
	"net/http"
	"net/http/pprof"
)

var R = render.New(render.Options{
	Layout:        "layout",
	Extensions:    []string{".html"},
	IsDevelopment: config.IsLocal(),
	Funcs:         []template.FuncMap{helpers.FuncMap},
})

func New(r *httprouter.Router) {
	r.GET("/", Top)
	r.GET("/en", TopEn)
	r.GET("/ja", TopJa)
	r.GET("/es", TopEs)
	r.POST("/", RedirectToTop)
	r.POST("/1/installations", handlers.LoginRequired(CreateInstallation))

	r.POST("/1/users", Signup)
	// This GET: /login endpoint is for backward compatibility.
	r.GET("/1/login", Login)
	r.POST("/1/login", PostLogin)
	r.POST("/1/fblogin", FBLogin)
	r.POST("/1/fblink", FBLink)
	r.POST("/1/twitterlogin", TwitterLogin)
	r.POST("/1/twitterLink", TwitterLink)
	r.POST("/1/applelogin", AppleLogin)
	r.POST("/1/applelink", AppleLink)
	r.POST("/1/logout", handlers.LoginRequired(Logout))
	r.GET("/1/users/:id", handlers.LoginRequired(GetUser))
	r.PUT("/1/users/:id", handlers.LoginRequired(UpdateUser))
	r.DELETE("/1/users/:id", handlers.LoginRequired(DeleteUser))
	r.PUT("/1/users/:id/email", handlers.LoginRequired(UpdateEmail))
	r.POST("/1/requestPasswordReset", RequestPasswordReset)
	r.POST("/1/confirmEmail", CreateConfirmEmail)
	r.PUT("/1/changePassword", handlers.LoginRequired(ChangePassword))
	r.GET("/1/referralProgram", handlers.LoginRequired(GetReferralProgram))
	r.GET("/1/badges", GetBadges)

	r.GET("/1/cities", handlers.LoginRequired(GetCities))
	r.GET("/1/sports", handlers.LoginRequired(GetSports))
	r.GET("/1/teamMemberApplication", handlers.LoginRequired(GetTeamMemberApplicationViewByUserId))
	r.GET("/1/teams", handlers.LoginRequired(GetTeams))
	r.GET("/1/demoteams", handlers.LoginRequired(GetDemoTeams))
	r.GET("/1/myteams", handlers.LoginRequired(GetMyTeams))
	r.GET("/1/teams/:teamId", handlers.LoginRequired(GetTeam))
	r.GET("/1/myteams/:teamId", handlers.LoginRequired(handlers.TeamMemberRequired(GetMyTeam)))
	r.GET("/1/followingTeams", handlers.LoginRequired(GetFollowingTeams))
	r.POST("/1/teams", handlers.LoginRequired(CreateTeam))
	r.PUT("/1/teams/:teamId", handlers.LoginRequired(handlers.TeamMemberRequired(UpdateTeam)))
	r.DELETE("/1/teams/:teamId", handlers.LoginRequired(handlers.TeamAdminRequired(DeleteTeam)))

	r.GET("/1/posts", handlers.LoginRequired(GetPosts))

	// 後方互換のため、c.coupon_type = 'codabar'のみcouponListを返却するendpoint
	r.GET("/1/teams/:teamId/coupons", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetCodabarCoupons))))

	r.GET("/1/teams/:teamId/posts", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamPosts))))
	r.GET("/1/teams/:teamId/posts/:postId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPost))))
	r.GET("/1/teams/:teamId/posts/:postId/comments", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPostComments))))
	r.GET("/1/teams/:teamId/posts/:postId/reactions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPostReactions))))
	r.GET("/1/teams/:teamId/posts/:postId/events", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPostEventAttendances))))
	r.GET("/1/teams/:teamId/posts/:postId/nonvoters", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetNonvoters))))
	r.GET("/1/teams/:teamId/posts/:postId/nonvotersByTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetNonvotersByTypes))))

	r.POST("/1/teams/:teamId/posts", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePost))))
	r.PUT("/1/teams/:teamId/posts/:postId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePost))))
	r.DELETE("/1/teams/:teamId/posts/:postId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeletePost))))

	r.POST("/1/teams/:teamId/posts/:postId/comments", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePostComment))))
	r.PUT("/1/teams/:teamId/posts/:postId/comments/:commentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePostComment))))
	r.DELETE("/1/teams/:teamId/posts/:postId/comments/:commentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeletePostComment))))

	r.POST("/1/teams/:teamId/posts/:postId/reactions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePostReaction))))

	r.GET("/1/teams/:teamId/posts/:postId/choices", handlers.LoginRequired(handlers.TeamMemberRequired(GetPostChoices)))
	r.GET("/1/teams/:teamId/posts/:postId/choices/:choiceId", handlers.LoginRequired(handlers.TeamMemberRequired(GetPostChoice)))
	r.POST("/1/teams/:teamId/posts/:postId/choices", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePostChoices))))
	r.DELETE("/1/teams/:teamId/posts/:postId/choices/:choiceId", handlers.LoginRequired(handlers.TeamMemberRequired(DeletePostChoice)))
	r.PUT("/1/teams/:teamId/posts/:postId/choices/:choiceId", handlers.LoginRequired(handlers.TeamMemberRequired(UpdatePostChoice)))
	r.POST("/1/teams/:teamId/posts/:postId/choices/:choiceId/votes", handlers.LoginRequired(handlers.TeamMemberRequired(CreatePostVote)))
	r.DELETE("/1/teams/:teamId/posts/:postId/choices/:choiceId/votes/:voteId", handlers.LoginRequired(handlers.TeamMemberRequired(DeletePostVote)))
	r.POST("/1/teams/:teamId/posts/:postId/reminders", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(RemindNonvoters))))

	r.GET("/1/teams/:teamId/teamMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMembers))))
	r.GET("/1/teams/:teamId/teamMembersByTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMembersByTypes))))
	r.GET("/1/teams/:teamId/teamMembersAndGuestsByTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMembersAndGuestsByTypes))))
	r.GET("/1/teams/:teamId/teamMembersWithAccountByTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMembersWithAccountByTypes))))
	r.GET("/1/teams/:teamId/teamMembersWithoutMyFamilyByTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMembersWithoutMyFamilyByTypes))))
	r.GET("/1/teams/:teamId/guestMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGuestMembers))))
	r.GET("/1/teams/:teamId/teamMembers/:teamMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMemberView))))
	r.GET("/1/teams/:teamId/users/:userId/teamMember", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMemberViewByUserId))))
	r.GET("/1/teams/:teamId/teamMembers/:teamMemberId/familyMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetFamilyMembers))))
	r.GET("/1/teams/:teamId/myFamilyMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetMyFamilyMembers))))
	r.GET("/1/teams/:teamId/scoreSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScoreSettings))))
	r.POST("/1/teams/:teamId/teamMembers", handlers.LoginRequired(handlers.TeamRequired(CreateTeamMember)))
	r.POST("/1/teams/:teamId/myFamilyMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(AddMyFamilyMembers))))
	r.PUT("/1/teams/:teamId/teamMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(UpdateTeamMembers))))
	r.PUT("/1/teams/:teamId/teamMembers/:teamMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateTeamMember))))
	r.PUT("/1/teams/:teamId/emailNotification", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateEmailNotification))))
	r.DELETE("/1/teams/:teamId/teamMembers/:teamMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteTeamMember))))
	r.DELETE("/1/teams/:teamId/myFamilyMembers/:teamMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteFamilyMember))))

	r.GET("/1/teams/:teamId/teamMemberApplications", handlers.LoginRequired(handlers.TeamAdminRequired(GetTeamMemberApplications)))
	r.GET("/1/teams/:teamId/teamMemberApplications/:teamMemberApplicationId", handlers.LoginRequired(handlers.TeamAdminRequired(GetTeamMemberApplicationView)))
	r.GET("/1/teams/:teamId/users/:userId/teamMemberApplication", handlers.LoginRequired(handlers.TeamAdminRequired(GetTeamMemberApplicationViewByTeamIdAndUserId)))
	r.POST("/1/teams/:teamId/teamMemberApplications", handlers.LoginRequired(handlers.DemoTeamForbidden(CreateTeamMemberApplication)))
	r.GET("/1/teams/:teamId/teamMemberApplications/:teamMemberApplicationId/approve", handlers.LoginRequired(handlers.TeamAdminRequired(ApproveTeamMemberApplicationAPI)))
	r.GET("/1/teams/:teamId/teamMemberApplications/:teamMemberApplicationId/disapprove", handlers.LoginRequired(handlers.TeamAdminRequired(DisapproveTeamMemberApplicationAPI)))

	r.GET("/1/teams/:teamId/followers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetFollowers))))
	r.GET("/1/teams/:teamId/followers/:followerId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetFollowerView))))
	r.PUT("/1/teams/:teamId/followers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(InviteFollowers))))
	r.DELETE("/1/teams/:teamId/followers/:followerId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(DeleteFollower))))
	r.GET("/1/teams/:teamId/followerApplications", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetFollowerApplications))))
	r.GET("/1/teams/:teamId/followerApplications/:followerApplicationId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetFollowerApplication))))
	r.GET("/1/teams/:teamId/followerApplications/:followerApplicationId/approve", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(ApproveFollowerApplicationAPI))))
	r.GET("/1/teams/:teamId/followerApplications/:followerApplicationId/disapprove", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(DisapproveFollowerApplicationAPI))))

	r.GET("/1/teams/:teamId/locations", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetLocations))))
	r.GET("/1/teams/:teamId/places", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPlaces))))
	r.GET("/1/teams/:teamId/placeDetails", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPlaceDetails))))
	r.GET("/1/teams/:teamId/locations/:locationId", handlers.LoginRequired(handlers.TeamMemberRequired(GetLocation)))
	r.POST("/1/teams/:teamId/locations", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateLocation))))
	r.PUT("/1/teams/:teamId/locations/:locationId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateLocation))))
	r.DELETE("/1/teams/:teamId/locations/:locationId", handlers.LoginRequired(handlers.TeamMemberRequired(DeleteLocation)))

	r.GET("/1/teams/:teamId/opponents", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetOpponents))))
	r.GET("/1/teams/:teamId/opponents/:opponentId", handlers.LoginRequired(handlers.TeamMemberRequired(GetOpponent)))
	r.POST("/1/teams/:teamId/opponents", handlers.LoginRequired(handlers.TeamMemberRequired(CreateOpponent)))
	r.PUT("/1/teams/:teamId/opponents/:opponentId", handlers.LoginRequired(handlers.TeamMemberRequired(UpdateOpponent)))
	r.DELETE("/1/teams/:teamId/opponents/:opponentId", handlers.LoginRequired(handlers.TeamMemberRequired(DeleteOpponent)))

	r.GET("/1/teams/:teamId/opponentMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetIndividualOpponentMembers))))
	r.GET("/1/teams/:teamId/opponents/:opponentId/opponentMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetOpponentMembers))))
	r.GET("/1/teams/:teamId/opponents/:opponentId/opponentMembers/:opponentMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetOpponentMember))))
	r.POST("/1/teams/:teamId/opponentMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateIndividualOpponentMember))))
	r.POST("/1/teams/:teamId/opponents/:opponentId/opponentMembers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateOpponentMember))))
	r.PUT("/1/teams/:teamId/opponents/:opponentId/opponentMembers/:opponentMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateOpponentMember))))
	r.DELETE("/1/teams/:teamId/opponents/:opponentId/opponentMembers/:opponentMemberId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteOpponentMember))))

	r.GET("/1/teams/:teamId/seasons", handlers.LoginRequired(handlers.TeamMemberRequired(GetSeasons)))
	r.GET("/1/teams/:teamId/seasons/:seasonId", handlers.LoginRequired(handlers.TeamMemberRequired(GetSeason)))
	r.POST("/1/teams/:teamId/seasons", handlers.LoginRequired(handlers.TeamMemberRequired(CreateSeason)))
	r.PUT("/1/teams/:teamId/seasons/:seasonId", handlers.LoginRequired(handlers.TeamMemberRequired(UpdateSeason)))
	r.DELETE("/1/teams/:teamId/seasons/:seasonId", handlers.LoginRequired(handlers.TeamMemberRequired(DeleteSeason)))

	r.POST("/1/teams/:teamId/subscriptions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(SubscribeToPlan))))
	r.PUT("/1/teams/:teamId/subscriptions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(UpdateCreditCard))))
	r.DELETE("/1/teams/:teamId/subscriptions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(CancelSubscription))))
	r.GET("/1/teams/:teamId/payments", handlers.LoginRequired(handlers.TeamAdminRequired(GetPayments)))
	r.POST("/1/teams/:teamId/inAppPurchases", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(InAppPurchaseToPlan))))
	r.GET("/1/teams/:teamId/mysubscriptions", handlers.LoginRequired(handlers.TeamAdminRequired(GetMySubscription)))

	r.GET("/1/teams/:teamId/tournaments", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetTournaments))))
	//このendpointはアプリとWeb版現在利用してないendpoint
	r.GET("/1/teams/:teamId/seasons/:seasonId/tournaments", handlers.LoginRequired(handlers.TeamMemberRequired(GetSeasonTournaments)))
	r.GET("/1/teams/:teamId/tournaments/:tournamentId", handlers.LoginRequired(handlers.TeamMemberRequired(GetTournament)))
	r.POST("/1/teams/:teamId/tournaments", handlers.LoginRequired(handlers.TeamMemberRequired(CreateTournament)))
	r.PUT("/1/teams/:teamId/tournaments/:tournamentId", handlers.LoginRequired(handlers.TeamMemberRequired(UpdateTournament)))
	r.DELETE("/1/teams/:teamId/tournaments/:tournamentId", handlers.LoginRequired(handlers.TeamMemberRequired(DeleteTournament)))

	r.GET("/1/teams/:teamId/equipments", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetEquipments))))
	r.GET("/1/teams/:teamId/equipments/:equipmentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetEquipment))))
	r.POST("/1/teams/:teamId/equipments", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateEquipment))))
	r.PUT("/1/teams/:teamId/equipments/:equipmentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateEquipment))))
	r.DELETE("/1/teams/:teamId/equipments/:equipmentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteEquipment))))

	r.GET("/1/teams/:teamId/events", handlers.LoginRequired(handlers.TeamMemberRequired(GetEvents)))
	r.GET("/1/teams/:teamId/upcomingEvents", handlers.LoginRequired(handlers.TeamMemberRequired(GetUpcomingEvents)))
	r.GET("/1/teams/:teamId/unrespondedEvents", handlers.LoginRequired(handlers.TeamMemberRequired(GetUnrespondedEvents)))
	r.GET("/1/teams/:teamId/pastEvents", handlers.LoginRequired(handlers.TeamMemberRequired(GetPastEvents)))
	r.GET("/1/teams/:teamId/gamesWithoutScores", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGamesWithoutScores))))
	r.GET("/1/teams/:teamId/seasons/:seasonId/gamesWithoutScores", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGamesWithoutScoresBySeason))))
	r.GET("/1/teams/:teamId/events/:eventId", handlers.LoginRequired(handlers.TeamMemberRequired(GetEvent)))
	r.GET("/1/teams/:teamId/events/:eventId/comments", handlers.LoginRequired(handlers.TeamMemberRequired(GetEventComments)))
	r.GET("/1/teams/:teamId/events/:eventId/attendance", handlers.LoginRequired(handlers.TeamMemberRequired(GetAttendance)))
	r.GET("/1/teams/:teamId/events/:eventId/myFamilyAttendances", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetMyFamilyAttendances))))
	r.GET("/1/teams/:teamId/events/:eventId/attendees", handlers.LoginRequired(handlers.TeamMemberRequired(GetAttendees)))
	r.GET("/1/teams/:teamId/events/:eventId/attendeesByTypes", handlers.LoginRequired(handlers.TeamMemberRequired(GetAttendeesByTypes)))
	r.GET("/1/teams/:teamId/events/:eventId/absentees", handlers.LoginRequired(handlers.TeamMemberRequired(GetAbsentees)))
	r.GET("/1/teams/:teamId/events/:eventId/absenteesByTypes", handlers.LoginRequired(handlers.TeamMemberRequired(GetAbsenteesByTypes)))
	r.GET("/1/teams/:teamId/events/:eventId/nonrespondents", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetNonrespondents))))
	r.GET("/1/teams/:teamId/events/:eventId/nonrespondentsByTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetNonrespondentsByTypes))))
	r.GET("/1/teams/:teamId/events/:eventId/rosters", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetRosters))))
	r.GET("/1/teams/:teamId/events/:eventId/score", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScoreByEventId))))
	r.GET("/1/teams/:teamId/scores", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScores))))
	r.GET("/1/teams/:teamId/seasons/:seasonId/tournaments/:tournamentId/scores", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScoresBySeasonAndTournament))))
	r.GET("/1/teams/:teamId/scores/:scoreId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScore))))
	r.GET("/1/teams/:teamId/events/:eventId/players", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPlayers))))
	r.GET("/1/teams/:teamId/events/:eventId/opponentPlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetOpponentPlayers))))
	r.GET("/1/teams/:teamId/events/:eventId/formations", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetFormations))))
	r.GET("/1/teams/:teamId/events/:eventId/ready", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetReady))))
	r.GET("/1/teams/:teamId/events/:eventId/notReady", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetNotReady))))
	r.GET("/1/teams/:teamId/events/:eventId/unconfirmed", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetUnconfirmed))))
	r.GET("/1/teams/:teamId/events/:eventId/gameSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGameSetting))))

	r.GET("/1/teams/:teamId/events/:eventId/hourWeather", handlers.LoginRequired(handlers.TeamMemberRequired(GetHourWeather)))
	r.GET("/1/teams/:teamId/events/:eventId/dayWeather", handlers.LoginRequired(handlers.TeamMemberRequired(GetDayWeather)))
	r.POST("/1/teams/:teamId/dayWeathers", handlers.LoginRequired(handlers.TeamMemberRequired(GetDayWeathers)))

	// The following three endpoints are for backward compatibility.
	r.GET("/1/teams/:teamId/footballScores/:scoreId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScore))))
	r.GET("/1/teams/:teamId/futsalScores/:scoreId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScore))))
	r.GET("/1/teams/:teamId/womensLacrosseScores/:scoreId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetScore))))

	r.GET("/1/teams/:teamId/scores/:scoreId/game", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGameByScoreId))))

	// The following three endpoints are for backward compatibility.
	r.GET("/1/teams/:teamId/footballScores/:scoreId/game", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGameByScoreId))))
	r.GET("/1/teams/:teamId/futsalScores/:scoreId/game", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGameByScoreId))))
	r.GET("/1/teams/:teamId/womensLacrosseScores/:scoreId/game", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetGameByScoreId))))

	r.GET("/1/teams/:teamId/foulTypes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamRequired(GetFoulTypes))))

	// for backward compatibility
	r.GET("/1/womensLacrosseFoulTypes", handlers.LoginRequired(GetWomensLacrosseFoulTypes))

	r.POST("/1/teams/:teamId/events", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateEvent))))
	r.PUT("/1/teams/:teamId/events/:eventId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateEvent))))
	r.DELETE("/1/teams/:teamId/events/:eventId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteEvent))))

	r.POST("/1/teams/:teamId/events/:eventId/comments", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateEventComment))))
	r.POST("/1/teams/:teamId/events/:eventId/attendances", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateAttendance))))
	r.POST("/1/teams/:teamId/events/:eventId/points", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePoint))))
	r.POST("/1/teams/:teamId/events/:eventId/cards", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateCard))))
	r.POST("/1/teams/:teamId/events/:eventId/substitutions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateSubstitution))))
	r.POST("/1/teams/:teamId/events/:eventId/shots", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateShot))))
	r.POST("/1/teams/:teamId/events/:eventId/draws", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateDraw))))
	r.POST("/1/teams/:teamId/events/:eventId/faceOffs", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateFaceOff))))
	r.POST("/1/teams/:teamId/events/:eventId/groundBalls", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateGroundBall))))
	r.POST("/1/teams/:teamId/events/:eventId/possessions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePossession))))
	r.POST("/1/teams/:teamId/events/:eventId/intercepts", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateIntercept))))
	r.POST("/1/teams/:teamId/events/:eventId/clears", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateClear))))
	r.POST("/1/teams/:teamId/events/:eventId/fouls", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateFoul))))
	r.POST("/1/teams/:teamId/events/:eventId/goals", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateGoal))))
	r.POST("/1/teams/:teamId/events/:eventId/stats", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateGameStats))))
	r.POST("/1/teams/:teamId/events/:eventId/scores", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateScore))))
	r.POST("/1/teams/:teamId/events/:eventId/reminders", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(RemindNonrespondents))))
	r.POST("/1/teams/:teamId/events/:eventId/players", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePlayers))))
	r.POST("/1/teams/:teamId/events/:eventId/opponentPlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateOpponentPlayers))))
	r.POST("/1/teams/:teamId/events/:eventId/gameSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpsertGameSetting))))
	r.POST("/1/teams/:teamId/events/:eventId/matchups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateMatchup))))

	r.PUT("/1/teams/:teamId/events/:eventId/comments/:commentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateEventComment))))
	r.PUT("/1/teams/:teamId/events/:eventId/attendances/:attendanceId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateAttendance))))
	r.PUT("/1/teams/:teamId/events/:eventId/scores/:scoreId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateScore))))
	r.PUT("/1/teams/:teamId/events/:eventId/scores/:scoreId/lineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateLineups))))
	r.PUT("/1/teams/:teamId/events/:eventId/formations", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateFormations))))
	r.PUT("/1/teams/:teamId/events/:eventId/goals/:goalId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateGoal))))
	r.PUT("/1/teams/:teamId/events/:eventId/cards/:cardId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateCard))))
	r.PUT("/1/teams/:teamId/events/:eventId/substitutions/:substitutionId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateSubstitution))))
	r.PUT("/1/teams/:teamId/events/:eventId/shots/:shotId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateShot))))
	r.PUT("/1/teams/:teamId/events/:eventId/draws/:drawId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateDraw))))
	r.PUT("/1/teams/:teamId/events/:eventId/faceOffs/:faceOffId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateFaceOff))))
	r.PUT("/1/teams/:teamId/events/:eventId/groundBalls/:groundBallId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateGroundBall))))
	r.PUT("/1/teams/:teamId/events/:eventId/intercepts/:interceptId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateIntercept))))
	r.PUT("/1/teams/:teamId/events/:eventId/possessions/:possessionId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePossession))))
	r.PUT("/1/teams/:teamId/events/:eventId/clears/:clearId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateClear))))
	r.PUT("/1/teams/:teamId/events/:eventId/fouls/:foulId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateFoul))))
	r.PUT("/1/teams/:teamId/events/:eventId/points/:pointId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePoint))))
	r.PUT("/1/teams/:teamId/events/:eventId/stats/:statsId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateGameStats))))
	r.PUT("/1/teams/:teamId/events/:eventId/players", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePlayers))))
	r.PUT("/1/teams/:teamId/events/:eventId/opponentPlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateOpponentPlayers))))
	r.PUT("/1/teams/:teamId/events/:eventId/players/:playerId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePlayer))))
	r.PUT("/1/teams/:teamId/events/:eventId/gameSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpsertGameSetting))))
	r.PUT("/1/teams/:teamId/events/:eventId/matchups/:matchupId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateMatchup))))
	r.POST("/1/teams/:teamId/events/:eventId/scores/:scoreId/notify", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(NotifyScoreUpdated))))

	r.GET("/1/teams/:teamId/baseballSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballSetting))))
	r.POST("/1/teams/:teamId/baseballSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateBaseballSetting))))
	r.PUT("/1/teams/:teamId/baseballSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballSetting))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballGameSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballGameSetting))))
	r.POST("/1/teams/:teamId/events/:eventId/baseballGameSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateBaseballGameSetting))))
	r.PUT("/1/teams/:teamId/events/:eventId/baseballGameSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballGameSetting))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballLineups))))
	r.POST("/1/teams/:teamId/events/:eventId/baseballLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballLineups))))
	r.PUT("/1/teams/:teamId/events/:eventId/baseballLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballLineups))))
	r.GET("/1/teams/:teamId/events/:eventId/restoreBaseballLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(RestoreBaseballLineups))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballOpponentLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballOpponentLineups))))
	// Typo alert! Keep it for backward compatibility.
	r.POST("/1/teams/:teamId/events/:eventId/basebalOpponentLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballOpponentLineups))))
	// Correct endpoint
	r.POST("/1/teams/:teamId/events/:eventId/baseballOpponentLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballOpponentLineups))))
	r.PUT("/1/teams/:teamId/events/:eventId/baseballOpponentLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballOpponentLineups))))
	r.GET("/1/teams/:teamId/events/:eventId/restoreBaseballOpponentLineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(RestoreBaseballOpponentLineups))))
	r.PUT("/1/teams/:teamId/events/:eventId/baseballLineups/:lineupId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballPlayer))))
	r.PUT("/1/teams/:teamId/events/:eventId/boxScores", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBoxScores))))
	r.PUT("/1/teams/:teamId/events/:eventId/boxScoreBattings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballBoxScoreBattings))))
	r.PUT("/1/teams/:teamId/events/:eventId/boxScoreFieldings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballBoxScoreFieldings))))
	r.PUT("/1/teams/:teamId/events/:eventId/boxScorePitchings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballBoxScorePitchings))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballSubstitutes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballSubstitutes))))
	r.GET("/1/teams/:teamId/events/:eventId/plays/:playId/baseballSubstitutes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballSubstitutes))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballReplaceablePlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballReplaceablePlayers))))
	r.GET("/1/teams/:teamId/events/:eventId/plays/:playId/baseballReplaceablePlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballReplaceablePlayers))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballOpponentSubstitutes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballOpponentSubstitutes))))
	r.GET("/1/teams/:teamId/events/:eventId/plays/:playId/baseballOpponentSubstitutes", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballOpponentSubstitutes))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballOpponentReplaceablePlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballOpponentReplaceablePlayers))))
	r.GET("/1/teams/:teamId/events/:eventId/plays/:playId/baseballOpponentReplaceablePlayers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballOpponentReplaceablePlayers))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballLineups/:lineupId/positionChanges", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballPositionChanges))))
	r.GET("/1/teams/:teamId/events/:eventId/plays", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetPlays))))
	r.GET("/1/teams/:teamId/events/:eventId/lastPlay", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetLastPlay))))
	r.POST("/1/teams/:teamId/events/:eventId/plays", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreatePlay))))
	r.PUT("/1/teams/:teamId/events/:eventId/plays/:playId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdatePlay))))
	r.GET("/1/teams/:teamId/events/:eventId/plays/:playId/undo", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UndoPlay))))
	r.GET("/1/teams/:teamId/events/:eventId/plays/:playId/redo", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(RedoPlay))))
	r.POST("/1/teams/:teamId/events/:eventId/baseballPlayEndHalfInning", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(CreateBaseballPlayEndHalfInning))))
	r.PUT("/1/teams/:teamId/events/:eventId/plays/:playId/baseballPlayEndHalfInning", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateBaseballPlayEndHalfInning))))
	r.GET("/1/teams/:teamId/events/:eventId/baseballTiebreaker", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetBaseballTiebreaker))))
	r.POST("/1/teams/:teamId/events/:eventId/baseballTiebreaker", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(StartBaseballTiebreaker))))
	r.PUT("/1/teams/:teamId/events/:eventId/plays/:playId/baseballTiebreaker", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(StartBaseballTiebreaker))))

	r.DELETE("/1/teams/:teamId/events/:eventId/comments/:commentId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteEventComment))))
	r.DELETE("/1/teams/:teamId/events/:eventId/goals/:goalId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteGoal))))
	r.DELETE("/1/teams/:teamId/events/:eventId/points/:pointId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeletePoint))))
	r.DELETE("/1/teams/:teamId/events/:eventId/cards/:cardId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteCard))))
	r.DELETE("/1/teams/:teamId/events/:eventId/stats/:gameStatsId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteGameStats))))
	r.DELETE("/1/teams/:teamId/events/:eventId/substitutions/:substitutionId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteSubstitution))))
	r.DELETE("/1/teams/:teamId/events/:eventId/substitutions", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteSubstitutions))))
	r.DELETE("/1/teams/:teamId/events/:eventId/shots/:shotId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteShot))))
	r.DELETE("/1/teams/:teamId/events/:eventId/draws/:drawId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteDraw))))
	r.DELETE("/1/teams/:teamId/events/:eventId/faceOffs/:faceOffId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteFaceOff))))
	r.DELETE("/1/teams/:teamId/events/:eventId/possessions/:possessionId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeletePossession))))
	r.DELETE("/1/teams/:teamId/events/:eventId/groundBalls/:groundBallId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteGroundBall))))
	r.DELETE("/1/teams/:teamId/events/:eventId/intercepts/:interceptId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteIntercept))))
	r.DELETE("/1/teams/:teamId/events/:eventId/clears/:clearId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteClear))))
	r.DELETE("/1/teams/:teamId/events/:eventId/fouls/:foulId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteFoul))))
	r.DELETE("/1/teams/:teamId/events/:eventId/scores/:scoreId/lineups", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteLineups))))
	r.DELETE("/1/teams/:teamId/events/:eventId/score", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteScore))))
	r.DELETE("/1/teams/:teamId/events/:eventId/matchups/matchupId", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(DeleteMatchup))))

	r.GET("/1/teams/:teamId/totalStats", handlers.LoginRequired(handlers.TeamMemberRequired(GetTotalStats)))
	r.GET("/1/teams/:teamId/stats", handlers.LoginRequired(handlers.TeamMemberRequired(GetStats)))
	r.GET("/1/teams/:teamId/statsByTournaments", handlers.LoginRequired(handlers.TeamMemberRequired(GetStatsByTournaments)))

	r.GET("/1/notifications", handlers.LoginRequired(GetUserNotifications))
	r.GET("/1/teams/:teamId/notifications", handlers.LoginRequired(handlers.TeamMemberRequired(GetNotifications)))
	r.PUT("/1/teams/:teamId/notifications/:notificationId", handlers.LoginRequired(handlers.TeamMemberRequired(UpdateNotification)))
	r.PUT("/1/teams/:teamId/notifications", handlers.LoginRequired(handlers.TeamMemberRequired(UpdateNotifications)))

	r.GET("/1/todos", handlers.LoginRequired(GetUserTodos))
	r.GET("/1/teams/:teamId/todos", handlers.LoginRequired(handlers.TeamMemberRequired(GetTodos)))
	r.POST("/1/chats/:roomId", handlers.LoginRequired(UploadChatImage))

	r.GET("/1/teams/:teamId/reminderSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetReminderSetting))))
	r.PUT("/1/teams/:teamId/reminderSettings", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(UpdateReminderSetting))))

	// The following three endpoints are for backward compatibility.
	r.POST("/1/video", handlers.LoginRequired(GetVideoProxy))
	r.GET("/1/video", handlers.LoginRequired(GetVideoProxy))
	r.POST("/1/videoQuery", v2.GraphqlHandler())

	r.GET("/2", v2.PlaygroundHandler())
	r.POST("/2/query", v2.GraphqlHandler())

	r.GET("/confirmEmail/:token", ShowConfirmEmail)
	r.GET("/confirmUpdateEmail/:token", ShowConfirmUpdateEmail)
	r.GET("/confirmEmailFollower/:token", ShowConfirmEmailFollower)

	r.GET("/resetPassword/:token", ShowResetPassword)
	r.POST("/resetPassword/:token", ResetPassword)

	r.GET("/approve/:path", ApproveTeamMemberApplication)
	r.GET("/disapprove/:path", DisapproveTeamMemberApplication)

	r.GET("/approveFollower/:path", ApproveFollowerApplication)
	r.GET("/disapproveFollower/:path", DisapproveFollowerApplication)

	r.GET("/going/:path", GoingToEvent)
	r.GET("/notgoing/:path", NotGoingToEvent)

	r.GET("/poll/:path", ViewPoll)
	r.POST("/poll/:path", Vote)

	// Firebase related
	r.GET("/1/iid/info/:iidToken", GetAppInstanceInfo)
	r.POST("/1/teams/:teamId/fcms", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(SendFCM))))
	r.POST("/1/teams/:teamId/fcmsToUsers", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(SendFCMToUsers))))
	r.POST("/1/topics/:topic/users", handlers.LoginRequired(SubscribeUsersToTopic))
	r.DELETE("/1/topics/:topic/users", handlers.LoginRequired(UnsubscribeUsersFromTopic))

	// Media API related
	r.GET("/1/teams/:teamId/mediaShareUrl", handlers.LoginRequired(handlers.TeamRequired(GetMediaShareUrl)))

	r.POST("/inquiry", SendInquiry)

	r.GET("/health", Health)

	r.POST("/1/stripehook", StripeHookForTeamHub)
	r.POST("/1/stripehookForPLAY", StripeHookForPLAY)

	r.POST("/1/teams/:teamId/referral/:referralCode", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamAdminRequired(CreateReferral))))
	r.GET("/1/teams/:teamId/referrals", handlers.LoginRequired(handlers.TeamRequired(handlers.TeamMemberRequired(GetInvitees))))

	r.GET("/teams/:teamId/invite", handlers.TeamRequired(InviteToTeamHub))
	r.GET("/teams/:teamId/inviteToPLAY", handlers.TeamRequired(InviteToPLAY))

	// Twitter OAuth
	r.GET("/twitterLogin", ShowTwitterLogin)
	r.GET("/twitterCallback", TwitterCallback)

	r.GET("/terms.html", Terms)
	r.GET("/terms_basic.html", TermsBasic)
	r.GET("/terms_plus.html", TermsPlus)
	r.GET("/privacy.html", Privacy)

	// Public team website
	r.GET("/teams/:teamId/scores/:scoreId", handlers.PublicTeamRequired(Score))
	r.GET("/teams/:teamId/events/:eventId", handlers.PublicTeamRequired(Event))
	r.GET("/teams/:teamId", handlers.PublicTeamRequired(TeamHome))
	r.GET("/teams/:teamId/home", handlers.PublicTeamRequired(TeamHome))
	r.GET("/teams/:teamId/events", handlers.PublicTeamRequired(TeamEvents))
	r.GET("/teams/:teamId/eventsInDate/:date", handlers.PublicTeamRequired(EventsInDate))
	r.GET("/teams/:teamId/intro", handlers.PublicTeamRequired(TeamIntro))
	r.GET("/teams/:teamId/users/:userId/teamMember", handlers.PublicTeamRequired(TeamMemberByUserId))
	r.GET("/teams/:teamId/members/:memberId", handlers.PublicTeamRequired(TeamMember))
	r.GET("/teams/:teamId/stats", handlers.PublicTeamRequired(TeamStats))
	r.GET("/teams/:teamId/blog", handlers.PublicTeamRequired(TeamBlog))
	r.GET("/teams/:teamId/contact", handlers.PublicTeamRequired(AdminContact))

	r.POST("/teams/:teamId/contact", handlers.PublicTeamRequired(CreateAdminContact))

	// Public API
	r.GET("/v1/baseballTeams", handlers.AuthenticateAPI(BaseballTeamsAPI))
	r.GET("/v1/baseballScores", handlers.AuthenticateAPI(BaseballScoresAPI))
	r.GET("/v1/baseballPlayerLeaderboards", handlers.AuthenticateAPI(BaseballPlayerLeaderboardsAPI))
	r.GET("/v1/baseballPlayerLeaderboard", handlers.AuthenticateAPI(BaseballPlayerLeaderboardAPI))
	r.GET("/v1/baseballTeamLeaderboards", handlers.AuthenticateAPI(BaseballTeamLeaderboardsAPI))
	r.GET("/v1/baseballTeamLeaderboard", handlers.AuthenticateAPI(BaseballTeamLeaderboardAPI))
	r.GET("/v1/baseballParks", handlers.AuthenticateAPI(BaseballParksAPI))
	r.GET("/v1/baseballParks/:baseballParkId", handlers.AuthenticateAPI(BaseballParkAPI))
	r.GET("/v1/battingCenters", handlers.AuthenticateAPI(BattingCentersAPI))
	r.GET("/v1/battingCenters/:battingCenterId", handlers.AuthenticateAPI(BattingCenterAPI))
	r.POST("/v1/teams/:teamId/fcms", handlers.AuthenticateAPI(handlers.TeamRequired(SendFCM)))
	r.POST("/v1/teams/:teamId/fcmsToUsers", handlers.AuthenticateAPI(handlers.TeamRequired(SendFCMToUsers)))

	// OAuth by TeamHub (provider functionality)
	r.GET("/v1/oauth", handlers.InitOAuth(OAuth))
	r.POST("/v1/oauth", handlers.LoginByEmail(OAuthLogin))
	r.POST("/v1/oauth/authorize", handlers.LoginRequired(OAuthorize))
	r.POST("/v1/oauth/token", handlers.AuthenticateOAuthClient(OAuthAccessToken))
	r.GET("/v1/oauth/user", handlers.AuthenticateOAuthToken(OAuthUser))
	r.GET("/v1/oauth/signup", handlers.InitOAuth(OAuthSignup))
	r.POST("/v1/oauth/signup", BeginOAuthSignup)
	r.POST("/v1/oauth/completeSignup", CompleteOAuthSignup)
	r.GET("/v1/oauth/confirmEmail/:path", OAuthConfirmEmail)
	r.GET("/v1/oauth/cancel", OAuthCancel)
	r.GET("/v1/oauth/forgotPassword", OAuthForgotPassword)
	r.POST("/v1/oauth/resetPassword", OAuthResetPassword)

	// OAuth by third-party providers (provider & consumer functionalities)
	r.GET("/v1/oauth/login/:provider", handlers.BeginOAuth(OAuthLogin))
	r.GET("/v1/oauth/callback/:provider", handlers.CompleteOAuth(OAuthLogin))
	r.POST("/v1/oauth/signup/:provider", handlers.SignupOAuth(OAuthLogin))

	// Internal API using ID token authentication
	r.GET("/v1/token/myteams/:teamId", handlers.AuthenticateToken(handlers.TeamMemberRequired(GetMyTeam)))
	r.GET("/v1/token/teams/:teamId/teamMembers", handlers.AuthenticateToken(handlers.TeamRequired(handlers.TeamMemberRequired(GetTeamMembers))))
	r.GET("/v1/token/teams/:teamId/teamMember", handlers.AuthenticateToken(handlers.TeamRequired(handlers.TeamMemberRequired(GetMyTeamMemberView))))
	r.POST("/v1/token/teams/:teamId/fcms", handlers.AuthenticateToken(handlers.TeamRequired(handlers.TeamMemberRequired(SendFCM))))
	r.POST("/v1/token/teams/:teamId/fcmsToUsers", handlers.AuthenticateToken(handlers.TeamRequired(handlers.TeamMemberRequired(SendFCMToUsers))))

	r.GET("/v2", graphHanders.PlaygroundHandler())
	r.POST("/v2/query", graphHanders.GraphqlHandler())

	// Shopify Multipass login for Joynup Gift, TeamHub Shop and TeamHub Order
	r.GET("/v1/shopify", Shopify)

	// The private routes for investigating Golang system metrics
	// https://go.dev/blog/pprof
	// https://github.com/google/pprof/blob/master/doc/README.md
	r.GET("/debug/pprof", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Index(w, req)
	})
	// cpu
	r.GET("/debug/profile", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Profile(w, req)
	})
	r.GET("/debug/symbol", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Symbol(w, req)
	})
	r.GET("/debug/trace", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Trace(w, req)
	})
	// blocking on synchronization primitives
	r.GET("/debug/block", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Handler("block").ServeHTTP(w, req)
	})
	// goroutine
	r.GET("/debug/goroutine", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Handler("goroutine").ServeHTTP(w, req)
	})
	r.GET("/debug/allocs", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Handler("allocs").ServeHTTP(w, req)
	})
	// memory
	r.GET("/debug/heap", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		pprof.Handler("heap").ServeHTTP(w, req)
	})
}

func Top(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	locale := req.Context().Value("locale").(string)
	t, _ := helpers.ParseTemplate("index", locale)
	t.Execute(w, nil)
}

func TopEn(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func TopJa(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	t, _ := template.ParseFiles("templates/index_ja.html")
	t.Execute(w, nil)
}

func TopEs(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	t, _ := template.ParseFiles("templates/index_es.html")
	t.Execute(w, nil)
}

func RedirectToTop(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	http.Redirect(w, req, config.AppURL, http.StatusFound)
}

func SendInquiry(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	req.ParseForm()
	helpers.ExecuteInBackground(func() {
		models.SendInquiry(
			req.FormValue("app"),
			req.FormValue("category"),
			req.FormValue("subject"),
			req.FormValue("name"),
			req.FormValue("email"),
			req.FormValue("text"),
		)
	})

	locale := req.Context().Value("locale").(string)
	T, _ := helpers.I18nTfunc(locale)

	R.HTML(w, http.StatusOK, "message", map[string]interface{}{"Message": T("contactSent")})
}

func Terms(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	locale := req.Context().Value("locale").(string)
	t, _ := helpers.ParseTemplate("terms", locale)
	t.Execute(w, nil)
}

func TermsBasic(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	locale := req.Context().Value("locale").(string)
	t, _ := helpers.ParseTemplate("terms_basic", locale)
	t.Execute(w, nil)
}

func TermsPlus(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	locale := req.Context().Value("locale").(string)
	t, _ := helpers.ParseTemplate("terms_plus", locale)
	t.Execute(w, nil)
}

func Privacy(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	locale := req.Context().Value("locale").(string)
	t, _ := helpers.ParseTemplate("privacy", locale)
	t.Execute(w, nil)
}

func Shopify(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	user := req.Context().Value("user").(*models.User)
	qs := req.URL.Query()
	url := models.GetShopifyURL(user, qs.Get("return_to"))
	http.Redirect(w, req, url, http.StatusFound)
}
