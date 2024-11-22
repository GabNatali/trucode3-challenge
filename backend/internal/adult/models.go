package adult

type Adult struct {
	Age           int           `json:"age"`
	Workclass     Workclass     `json:"workclass" gorm:"type:workclass_enum"`
	Fnlwgt        int           `json:"fnlwgt"`
	Education     Education     `json:"education" gorm:"type:education_enum"`
	EducationNum  EducationNum  `json:"education_num" gorm:"type:INTEGER"`
	MaritalStatus MaritalStatus `json:"marital_status" gorm:"type:marital_status_enum"`
	Occupation    Occupation    `json:"occupation" gorm:"type:occupation_enum"`
	Relationship  Relationship  `json:"relationship" gorm:"type:relationship_enum"`
	Race          Race          `json:"race" gorm:"type:race_enum"`
	Sex           Sex           `json:"sex" gorm:"type:sex_enum"`
	CapitalGain   int           `json:"capital_gain"`
	CapitalLoss   int           `json:"capital_loss"`
	HoursPerWeek  int           `json:"hours_per_week"`
	NativeCountry Country       `json:"native_country" gorm:"type:country_enum"`
	Income        Income        `json:"income" gorm:"type:income_enum"`
}

type Workclass string

const (
	Private            Workclass = "Private"
	SelfEmployed       Workclass = "Self-emp-inc"
	SelfEmployedNotInc Workclass = "Self-emp-not-inc"
	FederalGov         Workclass = "Federal-gov"
	LocalGov           Workclass = "Local-gov"
	StateGov           Workclass = "State-gov"
	WithoutPay         Workclass = "Without-pay"
	NeverWorked        Workclass = "Never-worked"
)

type Education string

const (
	SeventhEighth   Education = "7th-8th"
	Doctorate       Education = "Doctorate"
	Bachelors       Education = "Bachelors"
	SomeCollege     Education = "Some-college"
	Eleventh        Education = "11th"
	HSGrad          Education = "HS-grad"
	ProfSchool      Education = "Prof-school"
	AssocAcademic   Education = "Assoc-acdm"
	AssocVocational Education = "Assoc-voc"
	Ninth           Education = "9th"
	Seventh         Education = "7th"
	Twelfth         Education = "12th"
	Masters         Education = "Masters"
	FirstFourth     Education = "1st-4th"
	Tenth           Education = "10th"
	FifthSixth      Education = "5th-6th"
	Preschool       Education = "Preschool"
)

type EducationNum int

const (
	LevelBachelors       EducationNum = 13
	LevelSomeCollege     EducationNum = 10
	Level11th            EducationNum = 7
	LevelHSGrad          EducationNum = 9
	LevelProfSchool      EducationNum = 15
	LevelAssocAcademic   EducationNum = 12
	LevelAssocVocational EducationNum = 11
	Level9th             EducationNum = 5
	Level7th             EducationNum = 3
	Level12th            EducationNum = 8
	LevelMasters         EducationNum = 14
	Level1st4th          EducationNum = 2
	Level10th            EducationNum = 6
	Level5th6th          EducationNum = 3
	LevelPreschool       EducationNum = 1
	Level7th8th          EducationNum = 4
	LevelDoctorate       EducationNum = 16
)

type MaritalStatus string

const (
	MarriedCivSpouse    MaritalStatus = "Married-civ-spouse"
	Divorced            MaritalStatus = "Divorced"
	NeverMarried        MaritalStatus = "Never-married"
	Separated           MaritalStatus = "Separated"
	Widowed             MaritalStatus = "Widowed"
	MarriedSpouseAbsent MaritalStatus = "Married-spouse-absent"
	MarriedAfSpouse     MaritalStatus = "Married-AF-spouse"
)

type Occupation string

const (
	TechSupport      Occupation = "Tech-support"
	CraftRepair      Occupation = "Craft-repair"
	OtherService     Occupation = "Other-service"
	Sales            Occupation = "Sales"
	ExecManagerial   Occupation = "Exec-managerial"
	ProfSpecialty    Occupation = "Prof-specialty"
	HandlersCleaners Occupation = "Handlers-cleaners"
	MachineOpInspct  Occupation = "Machine-op-inspct"
	AdmCl            Occupation = "Adm-clerical"
	FarmingFishing   Occupation = "Farming-fishing"
	TransportMoving  Occupation = "Transport-moving"
	PrivHouseServ    Occupation = "Priv-house-serv"
	ProtectiveServ   Occupation = "Protective-serv"
	ArmedForces      Occupation = "Armed-Forces"
)

type Relationship string

const (
	Wife          Relationship = "Wife"
	OwnChild      Relationship = "Own-child"
	Husband       Relationship = "Husband"
	NotInFamily   Relationship = "Not-in-family"
	OtherRelative Relationship = "Other-relative"
	Unmarried     Relationship = "Unmarried"
)

type Race string

const (
	White            Race = "White"
	Black            Race = "Black"
	AsianPacIslander Race = "Asian-Pac-Islander"
	AmerIndianEskimo Race = "Amer-Indian-Eskimo"
	Others           Race = "Other"
)

type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

type Country string

const (
	UnitedStates Country = "United-States"
	Cambodia     Country = "Cambodia"
	England      Country = "England"
	PuertoRico   Country = "Puerto-Rico"
	Canada       Country = "Canada"
	Germany      Country = "Germany"
	OutlyingUS   Country = "Outlying-US(Guam-USVI-etc)"
	India        Country = "India"
	Japan        Country = "Japan"
	Greece       Country = "Greece"
	South        Country = "South"
	China        Country = "China"
	Cuba         Country = "Cuba"
	Iran         Country = "Iran"
	Honduras     Country = "Honduras"
	Philippines  Country = "Philippines"
	Italy        Country = "Italy"
	Poland       Country = "Poland"
	Jamaica      Country = "Jamaica"
	Vietnam      Country = "Vietnam"
	Mexico       Country = "Mexico"
	Portugal     Country = "Portugal"
	Ireland      Country = "Ireland"
	France       Country = "France"
	Dominican    Country = "Dominican-Republic"
	Laos         Country = "Laos"
	Ecuador      Country = "Ecuador"
	Taiwan       Country = "Taiwan"
	Haiti        Country = "Haiti"
	Columbia     Country = "Columbia"
	Hungary      Country = "Hungary"
	Guatemala    Country = "Guatemala"
	Nicaragua    Country = "Nicaragua"
	Scotland     Country = "Scotland"
	Thailand     Country = "Thailand"
	Yugoslavia   Country = "Yugoslavia"
	ElSalvador   Country = "El-Salvador"
	Trinadad     Country = "Trinadad&Tobago"
	Peru         Country = "Peru"
	Hong         Country = "Hong"
	Holand       Country = "Holand-Netherlands"
)

type Income string

const (
	LessEqual50K Income = "<=50K"
	Greater50K   Income = ">50K"
)

var EducationToLevel = map[Education]EducationNum{
	Bachelors:       LevelBachelors,
	SomeCollege:     LevelSomeCollege,
	Eleventh:        Level11th,
	HSGrad:          LevelHSGrad,
	ProfSchool:      LevelProfSchool,
	AssocAcademic:   LevelAssocAcademic,
	AssocVocational: LevelAssocVocational,
	Ninth:           Level9th,
	Seventh:         Level7th,
	Twelfth:         Level12th,
	Masters:         LevelMasters,
	FirstFourth:     Level1st4th,
	Tenth:           Level10th,
	FifthSixth:      Level5th6th,
	Preschool:       LevelPreschool,
}
