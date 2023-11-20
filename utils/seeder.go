package util

import model "attendance-is/models"

func Seed() {
	classes := []model.Class{
		{Name: "TIA2021"},
		{Name: "TIB2021"},
		{Name: "TIC2021"},
	}
	DB.Create(&classes)

	students := []model.Student{
		{Nim: "43321001", Name: "ADI RIFTA DWI KURNIAWAN"},
		{Nim: "43321002", Name: "ADRIAN REYHAN PRANATA"},
		{Nim: "43321003", Name: "AHMAD ADI IRFANSYAH"},
		{Nim: "43321004", Name: "ALFINA NAILAL MUNA"},
		{Nim: "43321005", Name: "ANATASYA OKTA BERLIANA"},
		{Nim: "43321006", Name: "CINTANIA SALSABELLA BERLIANT"},
		{Nim: "43321007", Name: "DENY WISNU SAPUTRO SUKISNO"},
		{Nim: "43321008", Name: "DZAKIY FAIZAL ABDUH"},
		{Nim: "43321009", Name: "FAJAR FATONI PRATAMA"},
		{Nim: "43321010", Name: "FIKRI MAJID"},

		{Nim: "43321101", Name: "AFIF IHZA DARMAWAN"},
		{Nim: "43321102", Name: "AHMAD SYAHRUL FAZA"},
		{Nim: "43321103", Name: "AINNUR HANIF NUGRAHA"},
		{Nim: "43321104", Name: "ARDHILLA EKA WINDIARTI"},
		{Nim: "43321105", Name: "ARYANTO ANDRI SOBIRIN"},
		{Nim: "43321106", Name: "BHATINDEN SEJIARGA ERGUN GIORTAGMA"},
		{Nim: "43321107", Name: "DEDY SETIAWAN EKA ARIANTO"},
		{Nim: "43321108", Name: "DWI RATNA PUSPITA SARI"},
		{Nim: "43321109", Name: "FEBRI ADI SETIYAWAN"},
		{Nim: "43321110", Name: "FIDO JAHFAL PRAYOGA"},

		{Nim: "43321201", Name: "AFIF RAMZY BADRANI"},
		{Nim: "43321202", Name: "AJI DWI PRAKOSO"},
		{Nim: "43321203", Name: "ASHABUL KAHFI"},
		{Nim: "43321204", Name: "BINA SATRIA FAUZI"},
		{Nim: "43321205", Name: "BUKHARY AZRIELLOREZQA YUFAR"},
		{Nim: "43321206", Name: "EKA YULIANTO"},
		{Nim: "43321207", Name: "FAUQA JAHFAL ISKANDAR"},
		{Nim: "43321208", Name: ""},
		{Nim: "43321209", Name: "GELORAWAN SUSATYO JATIPAMUNGKAS"},
		{Nim: "43321210", Name: "HENDI AHMAD SHOLEHUDIN"},
	}

	DB.Create(&students)

	DB.Model(&classes[0]).Association("Students").Replace(students[0:10])
	DB.Model(&classes[1]).Association("Students").Replace(students[10:20])
	DB.Model(&classes[2]).Association("Students").Replace(students[20:30])

	lecturer := []model.Lecturer{
		{Nip: "197904262003122002", Name: "KURNIANINGSIH, S.T., M.T., Dr."},
		{Nip: "198407192019031008", Name: "KUWAT SANTOSO, M. KOM"},
		{Nip: "197403112000121001", Name: "MARDIYONO, S.Kom., M.Sc."},
		{Nip: "199001072019031020", Name: "MUHAMMAD IRWAN YANWARI. S.Kom., M.Eng."},
		{Nip: "199107302019031010", Name: "NURSENO BAYU AJI, S. Kom, M. Kom."},
		{Nip: "198504102014041002", Name: "PRAYITNO, S.ST., M.T."},
		{Nip: "196810252000121001", Name: "TRI RAHARJO YUDANTORO, S. Kom., M.Kom."},
		{Nip: "198703272019022012", Name: "WIKTASARI S.T., M.Kom"},
		{Nip: "199004112019031014", Name: "AFANDI NUR AZIZ THOHARI, S.T., M.Cs"},
		{Nip: "198605292019032009", Name: "AISYATUL KARIMA, S.Kom., MCS"},
		{Nip: "198810142019031007", Name: "AMRAN YOBIOKTABERA, M. KOM"},
		{Nip: "199202052019031009", Name: "ANGGA WAHYU WIBOWO S.Kom., M.Eng"},
		{Nip: "197711192008012013", Name: "IDHAWATI H., S.Kom., M.Kom."},
		{Nip: "198404202015041003", Name: "LILIEK TRIYONO, S.T., M.Kom"},
		{Nip: "196008221988031001", Name: "PARSUMO RAHARDJO, Drs., M.Kom."},
		{Nip: "199401272019032036", Name: "SIRLI FAHRIAH, S.Kom, M.Kom."},
		{Nip: "197501302001121001", Name: "SLAMET HANDOKO, S.Kom., M.Kom."},
		{Nip: "197101172003121001", Name: "SUKAMTO, S.Kom., M.T."},
		{Nip: "197704012005011001", Name: "WAHYU SULISTIYO, S.T., M.Kom."},
	}

	DB.Create(&lecturer)

	courses := []model.Course{
		{Code: "432-191-502", Name: "Desain dan Perancangan Sistem Informasi"},
		{Code: "432-191-503", Name: "Sistem Multimedia"},
		{Code: "432-191-504", Name: "Internet of Things (IoT)"},
		{Code: "432-191-505", Name: "Data Mining"},
		{Code: "432-191-506", Name: "Komputasi Awan"},
		{Code: "432-191-508", Name: "Keamanan, Kesehatan, dan Keselamatan Kerja"},
		{Code: "432-191-509", Name: "Interaksi Manusia dan Komputer"},
	}

	DB.Create(&courses)

	for _, element := range classes {
		DB.Model(&element).Association("Courses").Replace(courses[0:7])
	}
	// DB.Model(&classes[1]).Association("Courses").Replace(courses[7:14])
	// DB.Model(&classes[2]).Association("Courses").Replace(courses[14:21])
	// DB.Model(courses[0:7]).Association("Courses").Replace(classes[0])
	// DB.Model(courses[7:14]).Association("Courses").Replace(classes[1])
	// DB.Model(courses[14:21]).Association("Courses").Replace(classes[2])

	// DB.Model(&lecturer[13]).Association("Courses").Replace(courses)
}
