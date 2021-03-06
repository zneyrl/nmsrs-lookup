package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"strconv"

	"strings"

	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
	"github.com/emurmotol/nmsrs/model"
	"github.com/unrolled/render"
)

// func GetRegistrants(w http.ResponseWriter, r *http.Request) {
// 	

// 	query := db.Model(&model.Registrant{})
// 	query.Count(&count)
// 	page, err := strconv.Atoi(r.URL.Query().Get("page"))

// 	if err != nil {
// 		page = 1
// 	}

// 	pagination := &helper.Paginator{
// 		Page:     page,
// 		Limit:    limit,
// 		Count:    count,
// 		Interval: interval,
// 		QueryUrl: r.URL.Query(),
// 	}

// 	if page > pagination.PageCount() {
// 		pagination.Page = 1
// 	}
// 	registrants := []model.Registrant{}
// 	query.Offset(pagination.Offset()).Limit(limit).Find(&registrants)

// 	data := make(map[string]interface{})
// 	data["title"] = "Registrants"
// 	data["auth"] = model.Auth(r)
// 	data["registrants"] = registrants
// 	data["q"] = r.URL.Query().Get("q")
// 	data["pagination"] = helper.Pager{
// 		Markup:     template.HTML(pagination.String()),
// 		Count:      pagination.Count,
// 		StartIndex: pagination.StartIndex(),
// 		EndIndex:   pagination.EndIndex(),
// 	}
// 	flashAlert := helper.GetFlash(w, r, "alert")

// 	if flashAlert != nil {
// 		alert := flashAlert.(helper.Alert)
// 		data["alert"] = template.HTML(alert.String())
// 	}
// 	rd.HTML(w, http.StatusOK, "registrant/index", data)
// }

func CreateRegistrant(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	createRegistrantForm := helper.GetFlash(w, r, "createRegistrantForm")

	if createRegistrantForm != nil {
		data["createRegistrantForm"] = createRegistrantForm.(model.CreateRegistrantForm)
	}
	data["civilStats"] = model.CivilStats()
	data["sexes"] = model.Sexes()
	data["empStats"] = model.EmpStats()
	data["disabilities"] = model.Disabilities()
	data["title"] = "Create Registrant"
	data["auth"] = model.Auth(r)
	rd.HTML(w, http.StatusOK, "registrant/create", data, render.HTMLOptions{Layout: "layouts/wizard"})
}

func StoreRegistrant(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		panic(err)
	}
	photoFile, photoHeader, err := r.FormFile("personalInfoPhoto")

	if err != nil {
		if err != http.ErrMissingFile {
			panic(err)
		}
	}
	delete(r.PostForm, "personalInfoPhoto")
	createRegistrantForm := new(model.CreateRegistrantForm)

	if err := decoder.Decode(createRegistrantForm, r.PostForm); err != nil {
		panic(err)
	}
	createRegistrantForm.PersonalInfoPhotoFile = photoFile
	createRegistrantForm.PersonalInfoPhotoHeader = photoHeader

	if !createRegistrantForm.IsValid() {
		helper.SetFlash(w, r, "createRegistrantForm", createRegistrantForm)
		CreateRegistrant(w, r)
		return
	}
	hasPhoto := false

	if createRegistrantForm.PersonalInfoPhotoFile != nil {
		hasPhoto = true
	}

	registrant := &model.Registrant{
		RegisteredAt: helper.ShortDate(createRegistrantForm.RegisteredAt),
		IAccept:      createRegistrantForm.IAccept,
		PersonalInfo: &model.PersonalInfo{
			HasPhoto:   hasPhoto,
			FamilyName: strings.ToUpper(createRegistrantForm.PersonalInfoFamilyName),
			GivenName:  strings.ToUpper(createRegistrantForm.PersonalInfoGivenName),
			MiddleName: strings.ToUpper(createRegistrantForm.PersonalInfoMiddleName),
			Birthdate:  helper.ShortDate(createRegistrantForm.PersonalInfoBirthdate),
			Password:   strings.ToUpper(createRegistrantForm.PersonalInfoPassword),
		},
		BasicInfo: &model.BasicInfo{
			StSub:          strings.ToUpper(createRegistrantForm.BasicInfoStSub),
			Province:       model.ProvinceById(bson.ObjectIdHex(createRegistrantForm.BasicInfoProvinceHexId)),
			Barangay:       model.BarangayById(bson.ObjectIdHex(createRegistrantForm.BasicInfoBarangayHexId)),
			PlaceOfBirth:   strings.ToUpper(createRegistrantForm.BasicInfoPlaceOfBirth),
			CivilStat:      model.CivilStatById(bson.ObjectIdHex(createRegistrantForm.BasicInfoCivilStatHexId)),
			CivilStatOther: strings.ToUpper(createRegistrantForm.BasicInfoCivilStatOther),
			Sex:            model.SexById(bson.ObjectIdHex(createRegistrantForm.BasicInfoSexHexId)),
			Age:            createRegistrantForm.BasicInfoAge,
			Weight:         createRegistrantForm.BasicInfoWeight,
			LandlineNumber: createRegistrantForm.BasicInfoLandlineNumber,
			MobileNumber:   createRegistrantForm.BasicInfoMobileNumber,
			Email:          strings.ToLower(createRegistrantForm.BasicInfoEmail),
		},
		Employment: &model.Employment{
			IsActivelyLookingForWork: createRegistrantForm.EmpIsActivelyLookingForWork,
			PassportNumber:           createRegistrantForm.EmpPassportNumber,
		},
		OtherSkillOther: strings.ToUpper(createRegistrantForm.OtherSkillOther),
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpStatHexId) {
		registrant.Employment.Stat = model.EmpStatById(bson.ObjectIdHex(createRegistrantForm.EmpStatHexId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpPrefLocalLocHexId) {
		registrant.Employment.PrefLocalLoc = model.CityMunById(bson.ObjectIdHex(createRegistrantForm.EmpPrefLocalLocHexId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpPrefOverseasLocHexId) {
		registrant.Employment.PrefOverseasLoc = model.CountryById(bson.ObjectIdHex(createRegistrantForm.EmpPrefOverseasLocHexId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.BasicInfoReligionHexId) {
		registrant.BasicInfo.Religion = model.ReligionById(bson.ObjectIdHex(createRegistrantForm.BasicInfoReligionHexId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.BasicInfoCityMunHexId) {
		registrant.BasicInfo.CityMun = model.CityMunById(bson.ObjectIdHex(createRegistrantForm.BasicInfoCityMunHexId))
	}

	if createRegistrantForm.BasicInfoHeightInFeet != 0 && createRegistrantForm.BasicInfoHeightInInches != 0 {
		registrant.BasicInfo.Height = &model.Height{
			Feet:   createRegistrantForm.BasicInfoHeightInFeet,
			Inches: createRegistrantForm.BasicInfoHeightInInches,
		}
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpUnEmpStatHexId) {
		registrant.Employment.UnEmpStat = model.UnEmpStatById(bson.ObjectIdHex(createRegistrantForm.EmpUnEmpStatHexId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpTeminatedOverseasCountryHexId) {
		registrant.Employment.TeminatedOverseasCountry = model.CountryById(bson.ObjectIdHex(createRegistrantForm.EmpTeminatedOverseasCountryHexId))
	}

	if createRegistrantForm.EmpPassportNumberExpiryDate != "" {
		registrant.Employment.PassportNumberExpiryDate = helper.YearMonth(createRegistrantForm.EmpPassportNumberExpiryDate)
	}

	if createRegistrantForm.DisabIsDisabled {
		registrant.Disab = &model.Disab{
			IsDisabled: createRegistrantForm.DisabIsDisabled,
			Name:       model.DisabilityById(bson.ObjectIdHex(createRegistrantForm.DisabHexId)),
			Other:      strings.ToUpper(createRegistrantForm.DisabOther),
		}
	}

	if len(createRegistrantForm.LangHexIds) != 0 {
		for _, langHexId := range createRegistrantForm.LangHexIds {
			registrant.Langs = append(registrant.Langs, model.LanguageById(bson.ObjectIdHex(langHexId)))
		}
	}

	if len(createRegistrantForm.EmpPrefOccHexIds) != 0 {
		for _, empPrefOccHexId := range createRegistrantForm.EmpPrefOccHexIds {
			registrant.Employment.PrefOccs = append(registrant.Employment.PrefOccs, model.PositionById(bson.ObjectIdHex(empPrefOccHexId)))
		}
	}

	if len(createRegistrantForm.OtherSkillHexIds) != 0 {
		for _, otherSkillHexId := range createRegistrantForm.OtherSkillHexIds {
			registrant.OtherSkills = append(registrant.OtherSkills, model.OtherSkillById(bson.ObjectIdHex(otherSkillHexId)))
		}
	}
	formalEduArr := []model.FormalEduArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.FormalEduJson), &formalEduArr); err != nil {
		panic(err)
	}

	if len(formalEduArr) != 0 {
		for _, formalEduObj := range formalEduArr {
			formalEdu := &model.FormalEdu{
				HighestGradeCompleted: model.EduLevelById(bson.ObjectIdHex(formalEduObj.HighestGradeCompletedHexId)),
				CourseDegree:          model.CourseById(bson.ObjectIdHex(formalEduObj.CourseDegreeHexId)),
				SchoolUnivOther:       strings.ToUpper(formalEduObj.SchoolUnivOther),
				YearGrad:              helper.Year(strconv.Itoa(formalEduObj.YearGrad)),
				LastAttended:          helper.YearMonth(formalEduObj.LastAttended),
			}

			if bson.IsObjectIdHex(formalEduObj.SchoolUnivHexId) {
				formalEdu.SchoolUniv = model.SchoolById(bson.ObjectIdHex(formalEduObj.SchoolUnivHexId))
			}
			registrant.FormalEdus = append(registrant.FormalEdus, formalEdu)
		}
	}
	proLicenseArr := []model.ProLicenseArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.ProLicenseJson), &proLicenseArr); err != nil {
		panic(err)
	}

	if len(proLicenseArr) != 0 {
		for _, proLicenseObj := range proLicenseArr {
			registrant.ProLicenses = append(registrant.ProLicenses, &model.ProLicense{
				Title:      model.LicenseById(bson.ObjectIdHex(proLicenseObj.TitleHexId)),
				ExpiryDate: helper.YearMonth(proLicenseObj.ExpiryDate),
			})
		}
	}
	eligArr := []model.EligArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.EligJson), &eligArr); err != nil {
		panic(err)
	}

	if len(eligArr) != 0 {
		for _, eligObj := range eligArr {
			registrant.Eligs = append(registrant.Eligs, &model.Elig{
				Title:     model.EligibilityById(bson.ObjectIdHex(eligObj.TitleHexId)),
				YearTaken: helper.YearMonth(eligObj.YearTaken),
			})
		}
	}
	trainingArr := []model.TrainingArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.TrainingJson), &trainingArr); err != nil {
		panic(err)
	}

	if len(trainingArr) != 0 {
		for _, trainingObj := range trainingArr {
			training := &model.Training{
				Name:                strings.ToUpper(trainingObj.Name),
				SkillsAcquired:      strings.ToUpper(trainingObj.SkillsAcquired),
				PeriodOfTrainingExp: strings.ToUpper(trainingObj.PeriodOfTrainingExp),
				CertReceived:        strings.ToUpper(trainingObj.CertReceived),
				IssuingSchoolAgency: strings.ToUpper(trainingObj.IssuingSchoolAgency),
			}
			registrant.Trainings = append(registrant.Trainings, training)
		}
	}
	certArr := []model.CertArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.CertJson), &certArr); err != nil {
		panic(err)
	}

	if len(certArr) != 0 {
		for _, certObj := range certArr {
			cert := &model.Cert{
				Title:      model.CertificateById(bson.ObjectIdHex(certObj.TitleHexId)),
				Rating:     strings.ToUpper(certObj.Rating),
				IssuedBy:   strings.ToUpper(certObj.IssuedBy),
				DateIssued: helper.YearMonth(certObj.DateIssued),
			}
			registrant.Certs = append(registrant.Certs, cert)
		}
	}
	workExpArr := []model.WorkExpArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.WorkExpJson), &workExpArr); err != nil {
		panic(err)
	}

	if len(workExpArr) != 0 {
		for _, workExpObj := range workExpArr {
			registrant.WorkExps = append(registrant.WorkExps, &model.WorkExp{
				NameOfCompanyFirm:    strings.ToUpper(workExpObj.NameOfCompanyFirm),
				Address:              strings.ToUpper(workExpObj.Address),
				PositionHeld:         model.PositionById(bson.ObjectIdHex(workExpObj.PositionHeldHexId)),
				From:                 helper.YearMonth(workExpObj.From),
				To:                   helper.YearMonth(workExpObj.To),
				IsRelatedToFormalEdu: workExpObj.IsRelatedToFormalEdu,
			})
		}
	}
	registrant.Create()
	http.Redirect(w, r, "/registrants", http.StatusFound)
}

func RegistrantEmailTaken(w http.ResponseWriter, r *http.Request) {
	if taken := model.RegistrantEmailTaken(r.URL.Query().Get("basicInfoEmail")); taken {
		data := make(map[string]string)
		data["error"] = lang.Get("emailTaken")
		rd.JSON(w, http.StatusNotFound, data)
		return
	}
	w.WriteHeader(http.StatusOK)
}
