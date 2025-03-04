package config

// constantes
const (
	PostgresTimestampLayout = "YYYY-MM-DD HH:MI:SS"

	ProjectModeDevelopment = "debug"
	ProjectModeTest        = "test"
	ProjectModeProduction  = "release"

	InfoSplitter   = "-#"
	ErrorSplitter  = "->"
	ServiceSuccess = "Service call successfully accomplished"
	DefaultLimit   = "10"
	DefaultPage    = "1"
)

// Error Messages
const (
	MessageBadRequestError      = "Noto'g'ri ma'lumot kiritildi"
	MessageNotFoundError        = "Ma'lumot topilmadi"
	MessageInternalServerError  = "Serverda ichki xatolik"
	MessageUnauthorizedError    = "Ro'yxatdan o'tilmagan"
	MessageUnauthenticatedError = "Ma'lumotlar tasdiqlanmadi"
)

// Error message in tech lan
const (
	TechMessageUniqueConstraint           = "duplicate key value violates unique constraint"
	TechMessageUniqueConstraintCompatible = "Kiritilgan ma'lumot tizimda mavjud"

	TechMessageNoRows           = "no rows in result set"
	TechMessageNoRowsCompatible = "Ma'lumot topilmadi"

	TechMessageRelationConstraint           = "violates foreign key constraint"
	TechMessageRelationConstraintCompatible = "Xato ma'lumot yuborildi"

	TechMessageEOF           = "EOF"
	TechMessageEOFCompatible = "Ma'lumot noto'g'ri kiritildi"

	TechMessageIncorrectPassword           = "Parol noto"
	TechMessageIncorrectPasswordCompatible = "Parol noto'g'ri kiritildi"

	TechMessageInternalServerCompatible = "Serverda ichki xatolik yuzaga keldi"
)

const MaxFileSize = 20 << 20  // 20 MB
const MaxVideoSize = 20 << 20 // 20 MB
const MaxImageSize = 5 << 20  // 5 MB

// Allowed image extensions.
var AllowedImageExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
}

// Allowed video extensions.
var AllowedVideoExt = map[string]bool{
	".mp4": true,
	".mov": true,
	".avi": true,
	".wmv": true,
	".mkv": true,
}

// Allowed file extensions.
var AllowedFileExt = map[string]bool{
	".pdf":  true,
	".xlsx": true,
	".xls":  true,
	".doc":  true,
	".docx": true,
	".ppt":  true,
	".pptx": true,
	// Add other formats as needed.
}

// Define allowed origins for production
var AllowedOrigins = []string{
	"https://api.domain.com",
}
