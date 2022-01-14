package ormodel

import (
	"gorm.io/gorm"
)

// ORProcess type for platform-dependent functionality
type ORProcess interface {

	// AcademicSessions
	HandleAddAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleDeleteAcademicSessions(orAcademicSessionsIDs []string) error
	HandleEditAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleAddOrEditAcademicSessions(orAcademicSessions []ORAcademicSessions) error

	// Users
	HandleAddUsers(orUser []ORUser, districtIDs []string) error
	HandleDeleteUsers(oruserIDs []string, districtIDs []string) error
	HandleEditUsers(orUser []ORUser) error
	HandleAddOrEditUsers(orUser []ORUser, districtIDs []string) error

	// Districts
	HandleAddDistrict(orOrg OROrg) (bool, error)
	HandleDeleteDistrict(orOrg OROrg) error
	HandleEditDistrict(orOrg OROrg, districtID string) error
	HandleAddOrEditDistrict(orOrg OROrg) error

	// Schools
	HandleAddSchool(orOrg OROrg, districtIDs []string) error
	HandleDeleteSchool(orOrg OROrg, districtIDs []string) error
	HandleEditSchool(orOrg OROrg) error
	HandleAddOrEditSchool(orOrg OROrg, districtIDs []string) error

	// Classes
	HandleAddClasses(orClass []ORClass, districtIDs []string) error
	HandleDeleteClasses(orClassIDs []string, districtIDs []string) error
	HandleEditClass(orClass []ORClass) error
	HandleAddOrEditClass(orClass []ORClass, districtIDs []string) error

	// Courses
	HandleAddCourses(orCourse []ORCourse, districtIDs []string) error
	HandleDeleteCourses(orCourseIDs []string, districtIDs []string) error
	HandleEditCourse(orCourse []ORCourse) error
	HandleAddOrEditCourse(orCourse []ORCourse, districtIDs []string) error

	// Enrollments
	HandleAddEnrollment(orEnrollment []OREnrollment, districtIDs []string) error
	HandleDeleteEnrollments(orEnrollment []OREnrollment, districtIDs []string) error
	HandleAddOrEditEnrollments(orEnrollment []OREnrollment, districtIDs []string) error

	RollBackOneRoster(orgDistrict []OROrg) error

	GetDistrictsIDs(orOrgs []OROrg) ([]string, error)
}

// OrManifest manifest file for oneroster
type OrManifest struct {
	gorm.Model
	PropertyName string `csv:"propertyName"`
	Value        string `csv:"value"`
}

// ORAcademicSessions academic sessions for oneroster
type ORAcademicSessions struct {
	gorm.Model
	SourcedID        string    `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                      //GUID
	Status           string    `csv:"status" json:"status" bson:"status,omitempty"`                               //Enumeration
	DateLastModified string    `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"` //DateTime
	Title            string    `csv:"title" json:"title" bson:"title,omitempty"`
	SessionType      string    `csv:"type" json:"type" bson:"type,omitempty"`                //Enumeration
	StartDate        string    `csv:"startDate" json:"startDate" bson:"startDate,omitempty"` //date
	EndDate          string    `csv:"endDate" json:"endDate" bson:"endDate,omitempty"`       //date
	ParentSourcedID  string    `csv:"parentSourcedId"`                                       //GUID Reference
	Parent           GUIDRef   `json:"parent" bson:"parent,omitempty"`
	Children         []GUIDRef `json:"children" bson:"children,omitempty"`
	SchoolYear       string    `csv:"schoolYear" json:"schoolYear" bson:"schoolYear,omitempty"` //year
}

// ORClass classes for oneroster
type ORClass struct {
	gorm.Model
	SourcedID        string    `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                      //GUID
	Status           string    `csv:"status" json:"status" bson:"status,omitempty"`                               //Enumeration
	DateLastModified string    `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"` //DateTime
	Title            string    `csv:"title" json:"title" bson:"title,omitempty"`
	Grades           string    `csv:"grades" json:"grades" bson:"grades,omitempty"` //[]string
	CourseSourcedID  string    `csv:"courseSourcedId"`                              //GUID Reference
	Course           GUIDRef   `json:"course" bson:"course,omitempty"`
	ClassCode        string    `csv:"classCode" json:"classCode" bson:"classCode,omitempty"`
	ClassType        string    `csv:"classType" json:"classType" bson:"classType,omitempty"` //Enumeration
	Location         string    `csv:"location" json:"location" bson:"location,omitempty"`
	SchoolSourcedID  string    `csv:"schoolSourcedId"` //GUID Reference
	School           GUIDRef   `json:"school" bson:"school,omitempty"`
	TermSourcedIds   string    `csv:"termSourcedIds" ` //List of GUID Reference
	Terms            []GUIDRef `json:"terms" bson:"terms,omitempty"`
	Subjects         string    `csv:"subjects" json:"subjects" bson:"subjects,omitempty"`
	SubjectCodes     string    `csv:"subjectCodes" json:"subjectCodes" bson:"subjectCodes,omitempty"`
	Periods          string    `csv:"periods" json:"periods" bson:"periods,omitempty"`
	Resources        []GUIDRef `json:"resources" bson:"resources,omitempty"`
}

// ORCourse courses for oneroster
type ORCourse struct {
	gorm.Model
	SourcedID           string  `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                      //GUID
	Status              string  `csv:"status" json:"status" bson:"status,omitempty"`                               //Enumeration
	DateLastModified    string  `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"` //DateTime
	SchoolYearSourcedID string  `csv:"schoolYearSourcedId"`                                                        //GUID Reference
	SchoolYear          GUIDRef `json:"schoolYear" bson:"schoolYear,omitempty"`
	Title               string  `csv:"title" json:"title" bson:"title,omitempty"`
	CourseCode          string  `csv:"courseCode" json:"courseCode" bson:"courseCode,omitempty"`
	// Grades				*[]string	`csv:"grades"`
	OrgSourcedID string  `csv:"orgSourcedId"` //GUID Reference
	Org          GUIDRef `json:"org" bson:"org,omitempty"`
	Subjects     string  `csv:"subjects" json:"subjects" bson:"subjects,omitempty"`
	SubjectCodes string  `csv:"subjectCodes" json:"subjectCodes" bson:"subjectCodes,omitempty"`
}

// ORDemographics demographics for one roster
type ORDemographics struct {
	gorm.Model
	SourcedID                            string `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                                                                                  //GUID
	Status                               string `csv:"status" json:"status" bson:"status,omitempty"`                                                                                           //Enumeration
	DateLastModified                     string `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"`                                                             //DateTime
	BirthDate                            string `csv:"birthDate" json:"birthDate" bson:"birthDate,omitempty"`                                                                                  //date
	Sex                                  string `csv:"sex" json:"sex" bson:"sex,omitempty"`                                                                                                    //Enumeration
	AmericanIndianOrAlaskaNative         string `csv:"americanIndianOrAlaskaNative" json:"americanIndianOrAlaskaNative" bson:"americanIndianOrAlaskaNative,omitempty"`                         //Enumeration
	Asian                                string `csv:"asian" json:"asian" bson:"asian,omitempty"`                                                                                              //Enumeration
	BlackOrAfricanAmerican               string `csv:"blackOrAfricanAmerican" json:"blackOrAfricanAmerican" bson:"blackOrAfricanAmerican,omitempty"`                                           //Enumeration
	NativeHawaiianOrOtherPacificIslander string `csv:"nativeHawaiianOrOtherPacificIslander" json:"nativeHawaiianOrOtherPacificIslander" bson:"nativeHawaiianOrOtherPacificIslander,omitempty"` //Enumeration
	White                                string `csv:"white" json:"white" bson:"white,omitempty"`                                                                                              //Enumeration
	DemographicRaceTwoOrMoreRaces        string `csv:"demographicRaceTwoOrMoreRaces" json:"demographicRaceTwoOrMoreRaces" bson:"demographicRaceTwoOrMoreRaces,omitempty"`                      //Enumeration
	HispanicOrLatinoEthnicity            string `csv:"hispanicOrLatinoEthnicity" json:"hispanicOrLatinoEthnicity" bson:"hispanicOrLatinoEthnicity,omitempty"`                                  //Enumeration
	CountryOfBirthCode                   string `csv:"countryOfBirthCode" json:"countryOfBirthCode" bson:"countryOfBirthCode,omitempty"`
	StateOfBirthAbbreviation             string `csv:"stateOfBirthAbbreviation" json:"stateOfBirthAbbreviation" bson:"stateOfBirthAbbreviation,omitempty"`
	CityOfBirth                          string `csv:"cityOfBirth" json:"cityOfBirth" bson:"cityOfBirth,omitempty"`
	PublicSchoolResidenceStatus          string `csv:"publicSchoolResidenceStatus" json:"publicSchoolResidenceStatus" bson:"publicSchoolResidenceStatus,omitempty"`
}

// OREnrollment enrollments for oneroster
type OREnrollment struct {
	gorm.Model
	SourcedID        string  `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                      //GUID
	Status           string  `csv:"status" json:"status" bson:"status,omitempty"`                               //Enumeration
	DateLastModified string  `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"` //DateTime
	ClassSourcedID   string  `csv:"classSourcedId"`                                                             //GUID Reference
	Class            GUIDRef `json:"class" bson:"class,omitempty"`
	SchoolSourcedID  string  `csv:"schoolSourcedId"` //GUID Reference
	School           GUIDRef `json:"school" bson:"school,omitempty"`
	UserSourcedID    string  `csv:"userSourcedId"` //GUID Reference
	User             GUIDRef `json:"user" bson:"user,omitempty"`
	Role             string  `csv:"role" json:"role" bson:"role,omitempty"` //Enumeration
	Primary          bool    `csv:"primary" json:"primary" bson:"primary,omitempty"`
	BeginDate        string  `csv:"beginDate" json:"beginDate" bson:"beginDate,omitempty"` //date
	EndDate          string  `csv:"endDate" json:"endDate" bson:"endDate,omitempty"`       //date
}

// OROrg orgs for oneroster
type OROrg struct {
	gorm.Model
	SourcedID        string    `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                      //GUID
	Status           string    `csv:"status" json:"status" bson:"status,omitempty"`                               //Enumeration
	DateLastModified string    `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"` //DateTime
	Name             string    `csv:"name" json:"name" bson:"name,omitempty"`
	OrgType          string    `csv:"type" json:"type" bson:"type,omitempty"` // type Enumeration
	Identifier       string    `csv:"identifier" json:"identifier" bson:"identifier,omitempty"`
	ParentSourcedID  string    `csv:"parentSourcedId"` //GUID Reference
	Parent           GUIDRef   `json:"parent" bson:"parent,omitempty"`
	Children         []GUIDRef `json:"children" bson:"children,omitempty"`
}

// ORUser users for oneroster
type ORUser struct {
	gorm.Model
	SourcedID        string          `csv:"sourcedId" json:"sourcedId" bson:"sourcedId,omitempty"`                      //GUID
	Status           string          `csv:"status" json:"status" bson:"status,omitempty"`                               //Enumeration
	DateLastModified string          `csv:"dateLastModified" json:"dateLastModified" bson:"dateLastModified,omitempty"` //DateTime
	EnabledUser      bool            `csv:"enabledUser" json:"enabledUser" bson:"enabledUser,omitempty"`
	OrgSourcedIds    string          `csv:"orgSourcedIds" json:"orgSourcedIds" bson:"orgSourcedIds,omitempty"` //List of GUID References.
	Orgs             []GUIDRef       `json:"orgs" bson:"orgs,omitempty"`
	Role             string          `csv:"role" json:"role" bson:"role,omitempty"` //Enumeration
	Username         string          `csv:"username" json:"username" bson:"username,omitempty"`
	UserIds          string          `csv:"userIds" json:"userIds" bson:"userIds,omitempty"` //[] string
	UserIdsIdentifer []UserIdentifer `json:"userIdsIdentifier" bson:"userIdsIdentifier,omitempty"`
	GivenName        string          `csv:"givenName" json:"givenName" bson:"givenName,omitempty"`
	FamilyName       string          `csv:"familyName" json:"familyName" bson:"familyName,omitempty"`
	MiddleName       string          `csv:"middleName" json:"middleName" bson:"middleName,omitempty"`
	Identifier       string          `csv:"identifier" json:"identifier" bson:"identifier,omitempty"`
	Email            string          `csv:"email" json:"email" bson:"email,omitempty"`
	Sms              string          `csv:"sms" json:"sms" bson:"sms,omitempty"`
	Phone            string          `csv:"phone" json:"phone" bson:"phone,omitempty"`
	AgentSourcedIds  string          `csv:"agentSourcedIds" json:"agentSourcedIds" bson:"agentSourcedIds,omitempty"` //List of GUID References
	Agents           []GUIDRef       `json:"agents" bson:"agents,omitempty"`
	Grades           string          `csv:"grades" json:"grades" bson:"grades,omitempty"`
	Password         string          `csv:"password" json:"password" bson:"password,omitempty"`
}

// ORCategory categories for oneroster
type ORCategory struct {
	gorm.Model
	SourcedID        string //GUID
	Status           string //Enumeration
	DateLastModified string //DateTime
	Title            string
}

// ORClassResources class resources for oneroster
type ORClassResources struct {
	gorm.Model
	SourcedID         string //GUID
	Status            string //Enumeration
	DateLastModified  string //DateTime
	Title             string
	ClassSourcedID    string //GUID Reference
	ResourceSourcedID string //GUID Reference
}

// ORCourseResources course resources for oneroster
type ORCourseResources struct {
	gorm.Model
	SourcedID         string //GUID
	Status            string //Enumeration
	DateLastModified  string //DateTime
	Title             string
	CourseSourcedID   string //GUID Reference
	ResourceSourcedID string //GUID Reference
}

// ORResource resource for oneroster
type ORResource struct {
	gorm.Model
	SourcedID        string //GUID
	Status           string //Enumeration
	DateLastModified string //DateTime
	VendorResourceID string //id
	Title            string
	Roles            []string //Enumeration List
	Importance       string
	VendorID         string //id
	ApplicationID    string //id
}

// ORResult results for oneroster
type ORResult struct {
	gorm.Model
	SourcedID         string  //GUID
	Status            string  //Enumeration
	DateLastModified  string  //DateTime
	LineItemSourcedID string  //GUID Reference
	StudentSourcedID  string  //GUID Reference
	ScoreStatus       string  //Enumeration
	Score             float64 //float
	ScoreDate         string  //date
	Comment           string
}

// ORLineItems line items for oneroster
type ORLineItems struct {
	gorm.Model
	SourcedID              string //GUID
	Status                 string //Enumeration
	DateLastModified       string //DateTime
	Title                  string
	Description            string
	SssignDate             string //date
	DueDate                string //date
	ClassSourcedID         string // GUID References
	CategorySourcedID      string // GUID References
	GradingPeriodSourcedID string // GUID References
	ResultValueMin         float64
	ResultValueMax         float64
}

// import type
const (
	ImportTypeBulk   = "bulk"
	ImportTypeDelta  = "delta"
	ImportTypeAbsent = "absent"
)

// manifest property names
const (
	ManifestProVersion              = "manifest.version"
	ManifestProOnerosterVersion     = "oneroster.version"
	ManifestProFileAcademicSessions = "file.academicSessions"
	ManifestProFileCategories       = "file.categories"
	ManifestProFileClasses          = "file.classes"
	ManifestProFileClassResources   = "file.classResources"
	ManifestProFileCourses          = "file.courses"
	ManifestProFileCourseResources  = "file.courseResources"
	ManifestProFileDemographics     = "file.demographics"
	ManifestProFileEnrollments      = "file.enrollments"
	ManifestProFileLineItems        = "file.lineItems"
	ManifestProFileOrgs             = "file.orgs"
	ManifestProFileResources        = "file.resources"
	ManifestProFileResults          = "file.results"
	ManifestProFileUsers            = "file.users"
	ManifestProSourceSystemName     = "source.systemName"
	ManifestProSourceSystemCode     = "source.systemCode"
)

// csv files name
const (
	CsvNameManifest         = "manifest.csv"
	CsvNameAcademicSessions = "academicSessions.csv"
	CsvNameCategories       = "categories.csv"
	CsvNameClasses          = "classes.csv"
	CsvNameCourses          = "courses.csv"
	CsvNameClassResources   = "classResources.csv"
	CsvNameDemographics     = "demographics.csv"
	CsvNameEnrollments      = "enrollments.csv"
	CsvNameOrgs             = "orgs.csv"
	CsvNameResources        = "resources.csv"
	CsvNameLineItems        = "lineItems.csv"
	CsvNameResults          = "results.csv"
	CsvNameUsers            = "users.csv"
)

// orgs types
const (
	OrgTypeDistrict = "district"
	OrgTypeSchool   = "school"
)

//Status types
const (
	StatusTypeActive      = "Active"
	StatusTypeToBeDeleted = "ToBeDeleted"
)

////// JSON ///////

// GUIDRef reference to GUID
type GUIDRef struct {
	gorm.Model
	Href      string `json:"href" bson:"href,omitempty"`
	SourcedID string `json:"sourcedId" bson:"sourcedId,omitempty"`
	GUIDType  string `json:"type" bson:"type,omitempty"`
}

// UserIdentifer holds identity for user
type UserIdentifer struct {
	gorm.Model
	Type       string `json:"type" bson:"type,omitempty"`
	Identifier string `json:"identifier" bson:"identifier,omitempty"`
}

// GUID constants
const (
	GUIDTypeAcademicSession = "academicSession"
	GUIDTypeCategory        = "category"
	GUIDTypeClass           = "class"
	GUIDTypeCourse          = "course"
	GUIDTypeDemographics    = "demographics"
	GUIDTypeEnrollment      = "enrollment"
	GUIDTypeOrg             = "org"
	GUIDTypeResource        = "resource"
	GUIDTypeLineItem        = "lineItem"
	GUIDTypeResult          = "result"
	GUIDTypeUser            = "user"
	GUIDTypeStudent         = "student"
	GUIDTypeTeacher         = "teacher"
	GUIDTypeTerm            = "term"
	GUIDTypeGradingPeriod   = "gradingPeriod"
)

//// Rest API Responses ////

// OrgsResponse for org responses
type OrgsResponse struct {
	Orgs []OROrg `json:"orgs" bson:"orgs,omitempty"`
}

// AcademicSessionsResponse for academic session responses
type AcademicSessionsResponse struct {
	AcademicSessions []ORAcademicSessions `json:"academicSessions" bson:"academicSessions,omitempty"`
}

// ClassesResponse for class responses
type ClassesResponse struct {
	Classes []ORClass `json:"classes" bson:"classes,omitempty"`
}

// CoursesResponse for course responses
type CoursesResponse struct {
	Courses []ORCourse `json:"courses" bson:"courses,omitempty"`
}

// EnrollmentsResponse for enrollment responses
type EnrollmentsResponse struct {
	Enrollments []OREnrollment `json:"enrollments" bson:"enrollments,omitempty"`
}

// UsersResponse for user responses
type UsersResponse struct {
	Users []ORUser `json:"users" bson:"users,omitempty"`
}
