package comicinfo

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function to generate a new ComicPageInfo for TESTING ONLY.
//
// This function will ensure XMLName is correctly assigned, with image & imageSize being specified.
// If usage is needed for other fields in ComicsPageInfo, Developer should use `ComicInfo{}` instead.
func newPage(image int, imageSize int64) ComicPageInfo {
	return ComicPageInfo{
		XMLName:   xml.Name{Space: "", Local: "Page"},
		Image:     image,
		ImageSize: imageSize,
	}
}

// cSpell:disable

// A comicInfo that generated from `resource/ComicInfo.xml`.
var testLoadResult = &ComicInfo{
	XMLName:         xml.Name{Local: "ComicInfo"},
	Xsi:             "http://www.w3.org/2001/XMLSchema-instance",
	Xsd:             "http://www.w3.org/2001/XMLSchema",
	Title:           "",
	Series:          "Say Hello to Blackjack",
	Number:          "1",
	Volume:          1,
	AlternateSeries: "",
	AlternateNumber: "",
	StoryArc:        "",
	SeriesGroup:     "",
	Summary: EscapedString{InnerXML: "Saitou Eijirou is a newly established intern doctor," +
		" who is forced to take on a second night job at another," +
		" much smaller hospital because of the extremely low pay he receives." +
		" As he bounces between the two different hospitals," +
		" he is forced to dig deeper and deeper into Japan's largely corrupted medical society" +
		" and starts to question even his own initial beliefs," +
		" as he asks himself just what being a doctor means."},
	Notes:           "",
	Writer:          "Shūhō Satō",
	Publisher:       "",
	Imprint:         "",
	Genre:           "Action, Drama, Romance, Medical",
	Web:             "",
	PageCount:       210,
	LanguageISO:     "en",
	Format:          "",
	Manga:           Manga_Yes,
	Characters:      "Eijiro Saito, Kaori Akagi, Kuniya Dekune, Yukiko Minagawa, Katsuo Ushida",
	Teams:           "",
	Locations:       "",
	ScanInformation: "v01",
	Pages: []ComicPageInfo{
		{XMLName: xml.Name{Space: "", Local: "Page"}, Image: 0, ImageSize: 1798991, Type: ComicPageType_FrontCover},
		newPage(1, 516297), newPage(2, 1436176), newPage(3, 1762430), newPage(4, 1901097), newPage(5, 1718912),
		newPage(6, 1076491), newPage(7, 1410195), newPage(8, 1405229), newPage(9, 1195708), newPage(10, 1394210),
		newPage(11, 1235988), newPage(12, 1056806), newPage(13, 1258274), newPage(14, 1185841), newPage(15, 1456285),
		newPage(16, 1391673), newPage(17, 1418986), newPage(18, 1161434), newPage(19, 1296058), newPage(20, 1459360),
		newPage(21, 1123603), newPage(22, 986919), newPage(23, 1213441), newPage(24, 1298574), newPage(25, 1017856),
		newPage(26, 1099335), newPage(27, 1050323), newPage(28, 1217915), newPage(29, 1193182), newPage(30, 1242906),
		newPage(31, 1310708), newPage(32, 1152725), newPage(33, 1346813), newPage(34, 1443641), newPage(35, 1599416),
		newPage(36, 1370287), newPage(37, 1326301), newPage(38, 1103030), newPage(39, 1447770), newPage(40, 1395890),
		newPage(41, 1616067), newPage(42, 1244963), newPage(43, 1288238), newPage(44, 1497089), newPage(45, 1092780),
		newPage(46, 1205222), newPage(47, 1267168), newPage(48, 1428456), newPage(49, 1351961), newPage(50, 1264990),
		newPage(51, 1319703), newPage(52, 1031108), newPage(53, 1183655), newPage(54, 1187989), newPage(55, 1319990),
		newPage(56, 1123313), newPage(57, 1005351), newPage(58, 38239), newPage(59, 38263), newPage(60, 1204980),
		newPage(61, 1458889), newPage(62, 1293875), newPage(63, 1410813), newPage(64, 1359816), newPage(65, 1204410),
		newPage(66, 1324269), newPage(67, 1433590), newPage(68, 1268696), newPage(69, 1048491), newPage(70, 1192405),
		newPage(71, 1401310), newPage(72, 1091939), newPage(73, 1316886), newPage(74, 1138429), newPage(75, 1122639),
		newPage(76, 1161792), newPage(77, 1141397), newPage(78, 969499), newPage(79, 1361513), newPage(80, 1139459),
		newPage(81, 922678), newPage(82, 1181528), newPage(83, 910314), newPage(84, 38239), newPage(85, 38263),
		newPage(86, 1140429), newPage(87, 1346022), newPage(88, 1119262), newPage(89, 872233), newPage(90, 1259365),
		newPage(91, 1327468), newPage(92, 1118956), newPage(93, 1026268), newPage(94, 1153095), newPage(95, 1342682),
		newPage(96, 1285151), newPage(97, 1181150), newPage(98, 1262304), newPage(99, 1392322), newPage(100, 1267421),
		newPage(101, 1439431), newPage(102, 1145011), newPage(103, 1419888), newPage(104, 758155), newPage(105, 1247489),
		newPage(106, 1372458), newPage(107, 1210961), newPage(108, 38239), newPage(109, 38263), newPage(110, 1700060),
		newPage(111, 1085854), newPage(112, 1513456), newPage(113, 1722329), newPage(114, 1100556), newPage(115, 1344053),
		newPage(116, 1144089), newPage(117, 1261382), newPage(118, 1413443), newPage(119, 1445150), newPage(120, 1155640),
		newPage(121, 1001424), newPage(122, 961317), newPage(123, 1237777), newPage(124, 1176664), newPage(125, 1201347),
		newPage(126, 1254286), newPage(127, 1479163), newPage(128, 1307920), newPage(129, 1429259), newPage(130, 1151533),
		newPage(131, 1343572), newPage(132, 1084292), newPage(133, 1182526), newPage(134, 1354181), newPage(135, 1331757),
		newPage(136, 1225946), newPage(137, 1131447), newPage(138, 1218902), newPage(139, 1195196), newPage(140, 1132278),
		newPage(141, 687091), newPage(142, 38239), newPage(143, 38263), newPage(144, 1420435), newPage(145, 1650666),
		newPage(146, 1120263), newPage(147, 1109990), newPage(148, 1296006), newPage(149, 1398963), newPage(150, 1237556),
		newPage(151, 1533610), newPage(152, 1462443), newPage(153, 1343406), newPage(154, 1157596), newPage(155, 1221897),
		newPage(156, 1385761), newPage(157, 1026578), newPage(158, 1212889), newPage(159, 1173074), newPage(160, 1090138),
		newPage(161, 1124369), newPage(162, 1239914), newPage(163, 1055504), newPage(164, 38239), newPage(165, 38263),
		newPage(166, 1155823), newPage(167, 738940), newPage(168, 1079359), newPage(169, 1279924), newPage(170, 1296372),
		newPage(171, 1418397), newPage(172, 1090120), newPage(173, 1324037), newPage(174, 1079965), newPage(175, 1180035),
		newPage(176, 1036674), newPage(177, 1208238), newPage(178, 1038531), newPage(179, 1142850), newPage(180, 1191809),
		newPage(181, 1269571), newPage(182, 1106513), newPage(183, 1223056), newPage(184, 1158605), newPage(185, 1006935),
		newPage(186, 38239), newPage(187, 38263), newPage(188, 1139596), newPage(189, 1486077), newPage(190, 1036622),
		newPage(191, 1074658), newPage(192, 1155241), newPage(193, 1254321), newPage(194, 1386042), newPage(195, 982884),
		newPage(196, 1204115), newPage(197, 1108878), newPage(198, 1272688), newPage(199, 1503069), newPage(200, 1365838),
		newPage(201, 1476189), newPage(202, 1229284), newPage(203, 1449817), newPage(204, 939341), newPage(205, 1329151),
		newPage(206, 1146394), newPage(207, 1030957), newPage(208, 38239), newPage(209, 638738),
	},
}

// cSpell:enable

func TestLoad(t *testing.T) {
	type testCase struct {
		path    string
		want    *ComicInfo
		wantErr bool
	}

	tests := []testCase{
		// 1. Graceful Load
		{"resources/ComicInfo.xml", testLoadResult, false},
		// 2. Missing "ComicInfo.xml" at path, but existed
		{"resources", nil, true},
		// 3. Pass invalid path
		{"", nil, true},
		// 4. Pass path that not exists
		{"resources/ComicInfo_not_exist.xml", nil, true},
	}

	for idx, tt := range tests {
		got, err := Load(tt.path)

		// Check error is wanted
		assert.Equalf(t, err != nil, tt.wantErr, "Case %d: Expected has any error: %v, but %v", idx, tt.wantErr, err)

		// Check comic info value
		assert.EqualValues(t, got, tt.want, "Case %d: values not equals", idx)
	}
}
