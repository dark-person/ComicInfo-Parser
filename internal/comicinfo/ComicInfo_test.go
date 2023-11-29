package comicinfo

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"testing"
)

// Test Parse a XML to ComicInfo Struct.
func TestParse(t *testing.T) {
	file, err := os.Open("Example.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := ComicInfo{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Check Result
	if v.Series != `Say Hello to Blackjack` || v.Number != "1" || v.Volume != 1 {
		t.Error("Wrong Value in Unmarshal")
	}

	if v.Summary.InnerXML != `Saitou Eijirou is a newly established intern doctor, who is forced to take on a second night job at another, much smaller hospital because of the extremely low pay he receives. As he bounces between the two different hospitals, he is forced to dig deeper and deeper into Japan's largely corrupted medical society and starts to question even his own initial beliefs, as he asks himself just what being a doctor means.` {
		t.Error("Wrong Summary in Unmarshal")
	}

	if v.Writer != `Shūhō Satō` {
		t.Error("Wrong Writer in Unmarshal")
	}

	if v.Genre != `Action, Drama, Romance, Medical` {
		t.Error("Wrong Genre in Unmarshal")
	}

	if v.PageCount != 210 || len(v.Pages) != 210 {
		t.Error("Wrong Page Count in Unmarshal")
	}

	if v.LanguageISO != "en" {
		t.Error("Wrong LanguageISO in Unmarshal")
	}

	if v.Manga != "Yes" {
		t.Errorf("Wrong Manga in Unmarshal")
	}

	if v.Characters != "Eijiro Saito, Kaori Akagi, Kuniya Dekune, Yukiko Minagawa, Katsuo Ushida" {
		t.Errorf("Wrong Character in Unmarshal")
	}

	if v.ScanInformation != "v01" {
		t.Errorf("Wrong ScanInformation in Unmarshal")
	}

}

// Test Marshal a XML to ComicInfo Struct, then Unmarshal to XML.
func TestMarshal(t *testing.T) {
	// Read Example XML in directory
	file, err := os.Open("Example.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Unmarshal XML
	v := ComicInfo{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Due to XML unknown problem
	v.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	v.Xsd = "http://www.w3.org/2001/XMLSchema"

	// Marshal XML
	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Check Result
	if string(output) != marshalResult {
		t.Errorf("Result not matched")
		os.Stdout.Write(output)
	}
}

// Test Create new ComicInfo object, then unmarshal to XML.
func TestNew(t *testing.T) {
	v := New()

	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	if string(output) != emptyMarshal {
		t.Errorf("Result not matched")
		os.Stdout.Write(output)
	}
}

// ===============
const marshalResult = `<ComicInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <Title></Title>
    <Series>Say Hello to Blackjack</Series>
    <Number>1</Number>
    <Volume>1</Volume>
    <AlternateSeries></AlternateSeries>
    <AlternateNumber></AlternateNumber>
    <StoryArc></StoryArc>
    <StoryArcNumber></StoryArcNumber>
    <SeriesGroup></SeriesGroup>
    <Summary>Saitou Eijirou is a newly established intern doctor, who is forced to take on a second night job at another, much smaller hospital because of the extremely low pay he receives. As he bounces between the two different hospitals, he is forced to dig deeper and deeper into Japan's largely corrupted medical society and starts to question even his own initial beliefs, as he asks himself just what being a doctor means.</Summary>
    <Notes></Notes>
    <Writer>Shūhō Satō</Writer>
    <Publisher></Publisher>
    <Imprint></Imprint>
    <Genre>Action, Drama, Romance, Medical</Genre>
    <Tags></Tags>
    <PageCount>210</PageCount>
    <LanguageISO>en</LanguageISO>
    <Format></Format>
    <AgeRating></AgeRating>
    <Manga>Yes</Manga>
    <Characters>Eijiro Saito, Kaori Akagi, Kuniya Dekune, Yukiko Minagawa, Katsuo Ushida</Characters>
    <Teams></Teams>
    <Locations></Locations>
    <ScanInformation>v01</ScanInformation>
    <Pages>
        <Page Image="0" Type="FrontCover" ImageSize="1798991"></Page>
        <Page Image="1" ImageSize="516297"></Page>
        <Page Image="2" ImageSize="1436176"></Page>
        <Page Image="3" ImageSize="1762430"></Page>
        <Page Image="4" ImageSize="1901097"></Page>
        <Page Image="5" ImageSize="1718912"></Page>
        <Page Image="6" ImageSize="1076491"></Page>
        <Page Image="7" ImageSize="1410195"></Page>
        <Page Image="8" ImageSize="1405229"></Page>
        <Page Image="9" ImageSize="1195708"></Page>
        <Page Image="10" ImageSize="1394210"></Page>
        <Page Image="11" ImageSize="1235988"></Page>
        <Page Image="12" ImageSize="1056806"></Page>
        <Page Image="13" ImageSize="1258274"></Page>
        <Page Image="14" ImageSize="1185841"></Page>
        <Page Image="15" ImageSize="1456285"></Page>
        <Page Image="16" ImageSize="1391673"></Page>
        <Page Image="17" ImageSize="1418986"></Page>
        <Page Image="18" ImageSize="1161434"></Page>
        <Page Image="19" ImageSize="1296058"></Page>
        <Page Image="20" ImageSize="1459360"></Page>
        <Page Image="21" ImageSize="1123603"></Page>
        <Page Image="22" ImageSize="986919"></Page>
        <Page Image="23" ImageSize="1213441"></Page>
        <Page Image="24" ImageSize="1298574"></Page>
        <Page Image="25" ImageSize="1017856"></Page>
        <Page Image="26" ImageSize="1099335"></Page>
        <Page Image="27" ImageSize="1050323"></Page>
        <Page Image="28" ImageSize="1217915"></Page>
        <Page Image="29" ImageSize="1193182"></Page>
        <Page Image="30" ImageSize="1242906"></Page>
        <Page Image="31" ImageSize="1310708"></Page>
        <Page Image="32" ImageSize="1152725"></Page>
        <Page Image="33" ImageSize="1346813"></Page>
        <Page Image="34" ImageSize="1443641"></Page>
        <Page Image="35" ImageSize="1599416"></Page>
        <Page Image="36" ImageSize="1370287"></Page>
        <Page Image="37" ImageSize="1326301"></Page>
        <Page Image="38" ImageSize="1103030"></Page>
        <Page Image="39" ImageSize="1447770"></Page>
        <Page Image="40" ImageSize="1395890"></Page>
        <Page Image="41" ImageSize="1616067"></Page>
        <Page Image="42" ImageSize="1244963"></Page>
        <Page Image="43" ImageSize="1288238"></Page>
        <Page Image="44" ImageSize="1497089"></Page>
        <Page Image="45" ImageSize="1092780"></Page>
        <Page Image="46" ImageSize="1205222"></Page>
        <Page Image="47" ImageSize="1267168"></Page>
        <Page Image="48" ImageSize="1428456"></Page>
        <Page Image="49" ImageSize="1351961"></Page>
        <Page Image="50" ImageSize="1264990"></Page>
        <Page Image="51" ImageSize="1319703"></Page>
        <Page Image="52" ImageSize="1031108"></Page>
        <Page Image="53" ImageSize="1183655"></Page>
        <Page Image="54" ImageSize="1187989"></Page>
        <Page Image="55" ImageSize="1319990"></Page>
        <Page Image="56" ImageSize="1123313"></Page>
        <Page Image="57" ImageSize="1005351"></Page>
        <Page Image="58" ImageSize="38239"></Page>
        <Page Image="59" ImageSize="38263"></Page>
        <Page Image="60" ImageSize="1204980"></Page>
        <Page Image="61" ImageSize="1458889"></Page>
        <Page Image="62" ImageSize="1293875"></Page>
        <Page Image="63" ImageSize="1410813"></Page>
        <Page Image="64" ImageSize="1359816"></Page>
        <Page Image="65" ImageSize="1204410"></Page>
        <Page Image="66" ImageSize="1324269"></Page>
        <Page Image="67" ImageSize="1433590"></Page>
        <Page Image="68" ImageSize="1268696"></Page>
        <Page Image="69" ImageSize="1048491"></Page>
        <Page Image="70" ImageSize="1192405"></Page>
        <Page Image="71" ImageSize="1401310"></Page>
        <Page Image="72" ImageSize="1091939"></Page>
        <Page Image="73" ImageSize="1316886"></Page>
        <Page Image="74" ImageSize="1138429"></Page>
        <Page Image="75" ImageSize="1122639"></Page>
        <Page Image="76" ImageSize="1161792"></Page>
        <Page Image="77" ImageSize="1141397"></Page>
        <Page Image="78" ImageSize="969499"></Page>
        <Page Image="79" ImageSize="1361513"></Page>
        <Page Image="80" ImageSize="1139459"></Page>
        <Page Image="81" ImageSize="922678"></Page>
        <Page Image="82" ImageSize="1181528"></Page>
        <Page Image="83" ImageSize="910314"></Page>
        <Page Image="84" ImageSize="38239"></Page>
        <Page Image="85" ImageSize="38263"></Page>
        <Page Image="86" ImageSize="1140429"></Page>
        <Page Image="87" ImageSize="1346022"></Page>
        <Page Image="88" ImageSize="1119262"></Page>
        <Page Image="89" ImageSize="872233"></Page>
        <Page Image="90" ImageSize="1259365"></Page>
        <Page Image="91" ImageSize="1327468"></Page>
        <Page Image="92" ImageSize="1118956"></Page>
        <Page Image="93" ImageSize="1026268"></Page>
        <Page Image="94" ImageSize="1153095"></Page>
        <Page Image="95" ImageSize="1342682"></Page>
        <Page Image="96" ImageSize="1285151"></Page>
        <Page Image="97" ImageSize="1181150"></Page>
        <Page Image="98" ImageSize="1262304"></Page>
        <Page Image="99" ImageSize="1392322"></Page>
        <Page Image="100" ImageSize="1267421"></Page>
        <Page Image="101" ImageSize="1439431"></Page>
        <Page Image="102" ImageSize="1145011"></Page>
        <Page Image="103" ImageSize="1419888"></Page>
        <Page Image="104" ImageSize="758155"></Page>
        <Page Image="105" ImageSize="1247489"></Page>
        <Page Image="106" ImageSize="1372458"></Page>
        <Page Image="107" ImageSize="1210961"></Page>
        <Page Image="108" ImageSize="38239"></Page>
        <Page Image="109" ImageSize="38263"></Page>
        <Page Image="110" ImageSize="1700060"></Page>
        <Page Image="111" ImageSize="1085854"></Page>
        <Page Image="112" ImageSize="1513456"></Page>
        <Page Image="113" ImageSize="1722329"></Page>
        <Page Image="114" ImageSize="1100556"></Page>
        <Page Image="115" ImageSize="1344053"></Page>
        <Page Image="116" ImageSize="1144089"></Page>
        <Page Image="117" ImageSize="1261382"></Page>
        <Page Image="118" ImageSize="1413443"></Page>
        <Page Image="119" ImageSize="1445150"></Page>
        <Page Image="120" ImageSize="1155640"></Page>
        <Page Image="121" ImageSize="1001424"></Page>
        <Page Image="122" ImageSize="961317"></Page>
        <Page Image="123" ImageSize="1237777"></Page>
        <Page Image="124" ImageSize="1176664"></Page>
        <Page Image="125" ImageSize="1201347"></Page>
        <Page Image="126" ImageSize="1254286"></Page>
        <Page Image="127" ImageSize="1479163"></Page>
        <Page Image="128" ImageSize="1307920"></Page>
        <Page Image="129" ImageSize="1429259"></Page>
        <Page Image="130" ImageSize="1151533"></Page>
        <Page Image="131" ImageSize="1343572"></Page>
        <Page Image="132" ImageSize="1084292"></Page>
        <Page Image="133" ImageSize="1182526"></Page>
        <Page Image="134" ImageSize="1354181"></Page>
        <Page Image="135" ImageSize="1331757"></Page>
        <Page Image="136" ImageSize="1225946"></Page>
        <Page Image="137" ImageSize="1131447"></Page>
        <Page Image="138" ImageSize="1218902"></Page>
        <Page Image="139" ImageSize="1195196"></Page>
        <Page Image="140" ImageSize="1132278"></Page>
        <Page Image="141" ImageSize="687091"></Page>
        <Page Image="142" ImageSize="38239"></Page>
        <Page Image="143" ImageSize="38263"></Page>
        <Page Image="144" ImageSize="1420435"></Page>
        <Page Image="145" ImageSize="1650666"></Page>
        <Page Image="146" ImageSize="1120263"></Page>
        <Page Image="147" ImageSize="1109990"></Page>
        <Page Image="148" ImageSize="1296006"></Page>
        <Page Image="149" ImageSize="1398963"></Page>
        <Page Image="150" ImageSize="1237556"></Page>
        <Page Image="151" ImageSize="1533610"></Page>
        <Page Image="152" ImageSize="1462443"></Page>
        <Page Image="153" ImageSize="1343406"></Page>
        <Page Image="154" ImageSize="1157596"></Page>
        <Page Image="155" ImageSize="1221897"></Page>
        <Page Image="156" ImageSize="1385761"></Page>
        <Page Image="157" ImageSize="1026578"></Page>
        <Page Image="158" ImageSize="1212889"></Page>
        <Page Image="159" ImageSize="1173074"></Page>
        <Page Image="160" ImageSize="1090138"></Page>
        <Page Image="161" ImageSize="1124369"></Page>
        <Page Image="162" ImageSize="1239914"></Page>
        <Page Image="163" ImageSize="1055504"></Page>
        <Page Image="164" ImageSize="38239"></Page>
        <Page Image="165" ImageSize="38263"></Page>
        <Page Image="166" ImageSize="1155823"></Page>
        <Page Image="167" ImageSize="738940"></Page>
        <Page Image="168" ImageSize="1079359"></Page>
        <Page Image="169" ImageSize="1279924"></Page>
        <Page Image="170" ImageSize="1296372"></Page>
        <Page Image="171" ImageSize="1418397"></Page>
        <Page Image="172" ImageSize="1090120"></Page>
        <Page Image="173" ImageSize="1324037"></Page>
        <Page Image="174" ImageSize="1079965"></Page>
        <Page Image="175" ImageSize="1180035"></Page>
        <Page Image="176" ImageSize="1036674"></Page>
        <Page Image="177" ImageSize="1208238"></Page>
        <Page Image="178" ImageSize="1038531"></Page>
        <Page Image="179" ImageSize="1142850"></Page>
        <Page Image="180" ImageSize="1191809"></Page>
        <Page Image="181" ImageSize="1269571"></Page>
        <Page Image="182" ImageSize="1106513"></Page>
        <Page Image="183" ImageSize="1223056"></Page>
        <Page Image="184" ImageSize="1158605"></Page>
        <Page Image="185" ImageSize="1006935"></Page>
        <Page Image="186" ImageSize="38239"></Page>
        <Page Image="187" ImageSize="38263"></Page>
        <Page Image="188" ImageSize="1139596"></Page>
        <Page Image="189" ImageSize="1486077"></Page>
        <Page Image="190" ImageSize="1036622"></Page>
        <Page Image="191" ImageSize="1074658"></Page>
        <Page Image="192" ImageSize="1155241"></Page>
        <Page Image="193" ImageSize="1254321"></Page>
        <Page Image="194" ImageSize="1386042"></Page>
        <Page Image="195" ImageSize="982884"></Page>
        <Page Image="196" ImageSize="1204115"></Page>
        <Page Image="197" ImageSize="1108878"></Page>
        <Page Image="198" ImageSize="1272688"></Page>
        <Page Image="199" ImageSize="1503069"></Page>
        <Page Image="200" ImageSize="1365838"></Page>
        <Page Image="201" ImageSize="1476189"></Page>
        <Page Image="202" ImageSize="1229284"></Page>
        <Page Image="203" ImageSize="1449817"></Page>
        <Page Image="204" ImageSize="939341"></Page>
        <Page Image="205" ImageSize="1329151"></Page>
        <Page Image="206" ImageSize="1146394"></Page>
        <Page Image="207" ImageSize="1030957"></Page>
        <Page Image="208" ImageSize="38239"></Page>
        <Page Image="209" ImageSize="638738"></Page>
    </Pages>
</ComicInfo>`

const emptyMarshal = `<ComicInfo xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <Title></Title>
  <Series></Series>
  <Number></Number>
  <Volume>0</Volume>
  <AlternateSeries></AlternateSeries>
  <AlternateNumber></AlternateNumber>
  <StoryArc></StoryArc>
  <StoryArcNumber></StoryArcNumber>
  <SeriesGroup></SeriesGroup>
  <Summary></Summary>
  <Notes></Notes>
  <Writer></Writer>
  <Publisher></Publisher>
  <Imprint></Imprint>
  <Genre></Genre>
  <Tags></Tags>
  <PageCount>0</PageCount>
  <LanguageISO></LanguageISO>
  <Format></Format>
  <AgeRating></AgeRating>
  <Manga></Manga>
  <Characters></Characters>
  <Teams></Teams>
  <Locations></Locations>
  <ScanInformation></ScanInformation>
  <Pages></Pages>
</ComicInfo>`
