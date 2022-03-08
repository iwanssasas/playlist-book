package errs

import "errors"

var (
	// sql error
	ErrDuplicateKey = errors.New("duplicate key")

	// common
	ErrForbidden         = errors.New("ErrForbidden")
	ErrUnauthorized      = errors.New("ErrUnauthorized")
	ErrPartnerNotAllowed = errors.New("ErrPartnerNotAllowed")
	ErrUserNotAllowed    = errors.New("cannot delete/edit this survey, because it was not created by You")

	// validate
	ErrIsRequired             = errors.New("is required")
	ErrPasswordDoestMatch     = errors.New("ErrPasswordDoestMatch")
	ErrTotalExposedIsNotEqual = errors.New("ErrTotalExposedIsNotEqual")
	ErrTotalReadyIsNotEqual   = errors.New("ErrTotalReadyIsNotEqual")
	ErrInvalidTimeBasis       = errors.New("ErrInvalidTimeBasis")
	ErrInvalidID              = errors.New("ErrInvalidID")
	ErrInvalidParams          = errors.New("ErrInvalidParams")
	ErrAtLeastZero            = errors.New("must at least 0")
	ErrMustBeFilled           = errors.New("must be filled")
	ErrInvalidContractor      = errors.New("ErrInvalidContractor")
	ErrAreaIsOccupied         = errors.New("Area is occupied during that time")
	ErrIsValidExcelType       = errors.New("IsValidExcelType")

	// geotech failure case
	ErrFailureCaseNotFound     = errors.New("ErrFailureCaseNotFound")
	ErrFailureCaseTypeNotFound = errors.New("ErrFailureCaseTypeNotFound")
	ErrFailureCaseNotDraft     = errors.New("ErrFailureCaseNotDraft")
	ErrFailureCaseIsDraft      = errors.New("ErrFailureCaseIsDraft")
	ErrFailureCaseIDNeeded     = errors.New("ErrFailureCaseIDNeeded")
	ErrFailureCaseNotRevised   = errors.New("ErrFailureCaseNotRevised")

	// geotech issue
	ErrIssueIDRequired           = errors.New("ErrIssueIDRequired")
	ErrUpdateIssueNeedID         = errors.New("ErrUpdateIssueNeedID")
	ErrFieldIsRequired           = errors.New("ErrFieldIsRequired")
	ErrMainPhotoIsRequired       = errors.New("ErrMainPhotoIsRequired")
	ErrDetailPhoto1IsRequired    = errors.New("ErrDetailPhoto1IsRequired")
	ErrDetailPhoto2IsRequired    = errors.New("ErrDetailPhoto2IsRequired")
	ErrIssueAttributesIsRequired = errors.New("ErrIssueAttributesIsRequired")
	ErrIssueCodeFailedToReset    = errors.New("ErrIssueCodeFailedToReset")
	ErrIssueNotDraft             = errors.New("ErrIssueNotDraft")
	ErrIssueNotRevised           = errors.New("ErrIssueNotRevised")
	ErrIssueNotFound             = errors.New("ErrIssueNotFound")

	// geotech monitoring
	ErrPatokHasInvalidValue      = errors.New("ErrPatokHasInvalidValue")
	ErrFirstPatokNotFound        = errors.New("ErrFirstPatokNotFound")
	ErrFailedCreateFirstPatok    = errors.New("ErrFailedCreateFirstPatok")
	ErrLastPatokNotFound         = errors.New("ErrLastPatokNotFound")
	ErrInvalidSequence           = errors.New("ErrInvalidSequence")
	ErrSkipAndReposInSameTime    = errors.New("cannot skip and repos a patok in one time")
	ErrMonitoringNotFound        = errors.New("Failed, Data Not Exist on DB")
	ErrMustBeAfterLatestDate     = errors.New("Failed, Patok date can't be the same / less than last patok")
	ErrMustBeAfterFirstPatokDate = errors.New("Failed, Patok date can't be the same / less than first patok")
	ErrPatokMustBeOnePatok       = errors.New("Failed, Patok must be minimal one Patok")
	ErrPatokSequenceNotFound     = errors.New("ErrPatokSequenceNotFound")
	ErrDontHaveFirstPatok        = errors.New("patok don't have first patok or previous patok")
	ErrCannotSkipAllPatoks       = errors.New("cannot skip all patoks, at least have 1 patok that don't skip")

	ErrPatokRepositionExist       = errors.New("Failed, Patok date can't be the same with patok reposition date")
	ErrMonitoringRepositionExist  = errors.New("Failed, Monitoring date can't be the same with patok reposition date")
	ErrDataMonitoringNotLastPatok = errors.New("Failed, Monitoring not last patok")

	// geotech weekly inspection
	ErrDateCannotBeFuture           = errors.New("ErrDateCannotBeFuture")
	ErrPercentageRange              = errors.New("ErrPercentageRange")
	ErrPercentageMustBeBiggerOrSame = errors.New("ErrPercentageMustBeBiggerOrSame")
	ErrInvalidImpactType            = errors.New("ErrInvalidImpactType")
	ErrInvalidLikelihood            = errors.New("ErrInvalidLikelihood")
	ErrInvalidRiskLevelType         = errors.New("ErrInvalidRiskLevelType")
	ErrInvalidStatusType            = errors.New("ErrInvalidStatusType")
	ErrInvalidPITWD                 = errors.New("ErrInvalidPITWD")

	// geotech geosurvey data
	ErrGeosurveyDataNotFound      = errors.New("ErrGeosurveyDataNotFound")
	ErrInvalidExplorationType     = errors.New("ErrInvalidExplorationType")
	ErrInvalidExplorationState    = errors.New("ErrInvalidExplorationState")
	ErrHasBeenThroughApprovalStep = errors.New("ErrHasBeenThroughApprovalStep")

	// Time
	ErrUploadTimeExeeded = errors.New("ErrUploadTimeExeeded")
	ErrDateCannotBeBlank = errors.New("ErrDateCannotBeBlank")
	ErrInvalidSchedule   = errors.New("ErrInvalidSchedule")
	ErrInvalidDate       = errors.New("ErrInvalidDate")
	ErrDateRange         = errors.New("date-from cannot be greater than date-to")

	// File
	ErrInvalidExtension  = errors.New("ErrInvalidExtension")
	ErrCsvFileNotValid   = errors.New("ErrCsvFileNotValid")
	ErrExcelFileNotValid = errors.New("ErrExcelFileNotValid")
	ErrFileMustPDF       = errors.New("ErrFileMustPDF")

	// not found
	ErrNotFound          = errors.New("ErrNotFound")
	ErrDataNotFound      = errors.New("data not found")
	ErrAreaNotFound      = errors.New("ErrAreaNotFound")
	ErrPartnerNotFound   = errors.New("ErrPartnerNotFound")
	ErrWeekNotFound      = errors.New("ErrWeekNotFound")
	ErrUserNotFound      = errors.New("ErrUserNotFound")
	ErrCompanyNotFound   = errors.New("ErrCompanyNotFound")
	ErrContractorNeeded  = errors.New("ErrContractorNotFound")
	ErrDataIsEmpty       = errors.New("ErrDataIsEmpty")
	ErrTokenNotExist     = errors.New("ErrTokenNotExist")
	ErrHaulDataNotExist  = errors.New("Error, you need to insert plan data first")
	ErrRadarNotFound     = errors.New("ErrRadarNotFound")
	ErrUserEmailNotFound = errors.New("ErrUserEmailNotFound")

	// required
	ErrImageRequired       = errors.New("ErrImageRequired")
	ErrFileRequired        = errors.New("ErrFileRequired")
	ErrIDRequired          = errors.New("ErrIDRequired")
	ErrCacheKeyRequired    = errors.New("ErrCacheKeyRequired")
	ErrUsernameRequired    = errors.New("ErrUsernameRequired")
	ErrUploadMapImage      = errors.New("ErrImageNeeded")
	ErrUploadMapContractor = errors.New("ErrContractorNeeded")
	ErrDateNeeded          = errors.New("ErrDateNeeded")
	ErrAreaNeeded          = errors.New("ErrAreaNeeded")
	ErrMonthNeeded         = errors.New("ErrMonthNeeded")
	ErrTypeNeeded          = errors.New("ErrTypeNeeded")
	ErrHourNeeded          = errors.New("ErrHourNeeded")
	ErrYearNeeded          = errors.New("ErrYearNeeded")
	ErrPartnerCodeRequired = errors.New("ErrPartnerCodeRequired")
	ErrPeriodIsRequired    = errors.New("ErrPeriodNeeded")

	// blasting
	ErrHasBeenApproved           = errors.New("ErrHasBeenApproved")
	ErrApprovalNeeded            = errors.New("ErrApprovalNeeded")
	ErrApprovalNotRevised        = errors.New("ErrApprovalNotRevised")
	ErrHasBeenRevised            = errors.New("ErrHasBeenRevised")
	ErrInvalidSheetName          = errors.New("Nama sheet nya salah, Nama sheet yang benar: INVENTORY")
	ErrNoApprovalBlastingData    = errors.New("There is no blasting request with fully approved")
	ErrAlreadyBlastingActivity   = errors.New("The blasting activity already exists")
	ErrTypeImageNotFound         = errors.New("Image type is not found")
	ErrTypeIdNotFound            = errors.New("Error type id is not found")
	ErrFileNotInputed            = errors.New("File not inputed")
	ErrIdTypeFilled              = errors.New("Error id and type is filled")
	ErrBlastingActivityNotFound  = errors.New("Blasting activity not found")
	ErrBlastingActivityLimitTime = errors.New("The blasting date exceeds the specified time")
	ErrWrongMaterial             = errors.New("invalid material")
	ErrMaximumRequestInOneDay    = errors.New("Error, the maximum request in one day is one request")

	ErrTypeNotValid       = errors.New("Type not valid")
	ErrMonthNotValid      = errors.New("Month not valid")
	ErrDateDailyNotValid  = errors.New("Error date daily not valid")
	ErrDateWeeklyNotValid = errors.New("Error date weekly not valid")
	ErrFieldIsWrongResult = errors.New("Error field is wrong")

	// mine design
	ErrImageAerialViewRequired = errors.New("ErrAerialViewImageRequired")
	ErrPeriodRequired          = errors.New("ErrPeriodRequired")
	ErrPartnerAreaRequired     = errors.New("ErrPartnerAreaRequired")
	ErrWorkOrderNotFound       = errors.New("Error Contractor in Work Order Not Found")

	// mine review
	ErrMaxSizeImage = errors.New("Image Size too Large")

	// production
	ErrInvalidActualType = errors.New("ErrInvalidActualType")
	ErrUnitRequired      = errors.New("ErrUnitRequired")
	ErrInvalidStartTime  = errors.New("ErrInvalidStartTime")
	ErrInvalidEndTime    = errors.New("ErrInvalidEndTime")
	ErrFailedRow         = errors.New("there is failed row")

	// coal-stock
	ErrInvalidCoalStockType = errors.New("ErrInvalidCoalStockType")
	ErrUploadExcelFailed    = errors.New("ErrUploadExcelFailed")

	// approval
	ErrInvalidWeek           = errors.New("ErrInvalidWeek")
	ErrReviseNotesIsRequired = errors.New("ErrReviseNotesIsRequired")

	// work-order
	ErrMaxSizeFile = errors.New("File Size too Large")
	ErrWoExist     = errors.New("Work Order is Already Exist")
	ErrDataExist   = errors.New("Data is already exist")

	// geofencing
	ErrGeofencingAlreadyExist = errors.New("Geofencing with that coordinate already exist")

	// Comment
	ErrorInputEmptyComment = errors.New("ErrorInputEmptyComment")
)
