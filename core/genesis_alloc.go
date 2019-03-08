// Copyright 2017 The go-simplechain Authors
// This file is part of the go-simplechain library.
//
// The go-simplechain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-simplechain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-simplechain library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const testnetAllocData = "\xe4\xe3\x94\t\xea\xf1Ak\xcd\xcbE\xa8\a\u06bf\xda\u00bd\\\xf7C\xa9\xe5\x8d\f\x9f,\x9c\xd0Ft\xed\xea@\x00\x00\x00"
const defaultGenesisBlockExtraData = "0x" +
	"54686520666972737420626c6f636b636861696e20617474657374617469" +
	"6f6e2061636b6e6f776c656467656420627920636f7572743a0a32382f30" +
	"362f3230313820e69dade5b79ee4ba92e88194e7bd91e6b395e999a2e987" +
	"87e7bab3e4ba86e59fbae4ba8ee4bf9de585a8e7bd9142616f5175616e2e" +
	"636f6de5ad98e582a8e79a84e585b3e994aee794b5e5ad90e8af81e68dae" +
	"e5b9b6e5819ae587bae58fb8e6b395e588a4e4be8befbc8ce58cbae59d97" +
	"e993bee794b5e5ad90e5ad98e8af81e79a84e6b395e5be8be69588e58a9b" +
	"e9a696e6aca1e88eb7e5be97e689bfe8aea4efbc8ce5ae83e6a087e5bf97" +
	"e79d80e58cbae59d97e993bee68a80e69cafe4b88ee4babae7b1bbe78eb0" +
	"e8a18ce7a4bee4bc9ae8a784e58899e79a84e89e8de59088efbc8ce4b8ba" +
	"e4bbb7e580bce4ba92e993bee7bd91e79a84e69e84e5bbbae5a5a0e5ae9a" +
	"e4ba86e6a0b9e59fbae380820a32382f30362f323031382048616e677a68" +
	"6f7520496e7465726e657420436f757274206d6164652074686520666972" +
	"73742073656e74656e6365206261736564206f6e2074686520656c656374" +
	"726f6e696320646174612065766964656e63652066726f6d2042616f5175" +
	"616e2e636f6d2e2054686520626c6f636b636861696e2061747465737461" +
	"74696f6e207468657265666f726520686173206265656e2061636b6e6f77" +
	"6c6564676564206279206a7564696369616c2073797374656d20666f7220" +
	"7468652066697273742074696d6520616e64206d65616e696e6766756c6c" +
	"792073796d626f6c6973656420746865206d657267696e67206f6620626c" +
	"6f636b636861696e20746563686e6f6c6f677920616e642063757272656e" +
	"742068756d616e20736f63696574792072756c65732e204974206c616964" +
	"2074686520666f756e646174696f6e20666f7220636f6e73747275637469" +
	"6e672074686520696e7465726e6574206f662076616c75652e0a0a546865" +
	"206f726967696e616c20466f756e646174696f6e20436f756e63696c3a0a" +
	"48616e73656e2047616f2c204c656f2059752c20436861647769636b204c" +
	"65652c205a68616f6c696e205969702c204b727a79737a746f6620506965" +
	"63680a546865206f726967696e616c20546563686e6963616c2053746565" +
	"72696e6720436f6d6d69747465653a204d6f726f205a68616e672c204a65" +
	"616e205a68616e670a0a546865204e65777320686561646c696e65206f66" +
	"20746865206461793a0a5554432b382030332f30312f323031392031303a" +
	"3236414d204368616e6727652034206d616b657320776f726c642d666972" +
	"7374206c616e64696e67206f6e206d6f6f6e27732066617220736964650a" +
	"4c756e617220496d70616374204372617465723a20566f6e204bc3a1726d" +
	"c3a16e0a436f6f7264696e617465733a2034352e35c2b053203137372e36" +
	"c2b0450a46697273742070686f746f206f6620746865206d6f6f6e277320" +
	"66617220736964653a0a646174613a696d6167652f6a7065673b62617365" +
	"36342c2f396a2f34514159525868705a67414153556b7141416741414141" +
	"41414141414141414141502f734142464564574e72655141424141514141" +
	"41414b4141442f34514e746148523063446f764c32357a4c6d466b62324a" +
	"6c4c6d4e7662533934595841764d5334774c7741385033687759574e725a" +
	"585167596d566e6157343949752b377679496761575139496c6331545442" +
	"4e63454e6c61476c49656e4a6c5533704f56474e3661324d355a43492f50" +
	"69413865447034625842745a585268494868746247357a4f6e6739496d46" +
	"6b62324a6c4f6d357a4f6d316c64474576496942344f6e68746348527250" +
	"534a425a4739695a53425954564167513239795a5341314c6a5974597a45" +
	"304e5341334f5334784e6a4d304f546b73494449774d5467764d4467764d" +
	"544d744d5459364e4441364d6a496749434167494341674943492b494478" +
	"795a475936556b5247494868746247357a4f6e4a6b5a6a30696148523063" +
	"446f764c336433647935334d793576636d63764d546b354f5338774d6938" +
	"794d6931795a47597463336c75644746344c57357a4979492b494478795a" +
	"4759365247567a59334a706348527062323467636d526d4f6d4669623356" +
	"3050534969494868746247357a4f6e68746345314e50534a6f644852774f" +
	"693876626e4d7559575276596d5575593239744c336868634338784c6a41" +
	"76625730764969423462577875637a707a64464a6c5a6a30696148523063" +
	"446f764c32357a4c6d466b62324a6c4c6d4e7662533934595841764d5334" +
	"774c334e556558426c4c314a6c63323931636d4e6c556d566d4979496765" +
	"473173626e4d366547317750534a6f644852774f693876626e4d75595752" +
	"76596d5575593239744c336868634338784c6a4176496942346258424e54" +
	"547050636d6c6e6157356862455276593356745a57353053555139496e68" +
	"746343356b615751364d6a45784e6a593251555977526a424752546b784d" +
	"5546464f4467354d5451314d7a49784d444d334f44596949486874634531" +
	"4e4f6b5276593356745a57353053555139496e68746343356b6157513651" +
	"55593552544a44526b45774e7a677a4d5446464f54677751304d354e4559" +
	"794e6b55774d7a51774e454969494868746345314e4f6b6c756333526862" +
	"6d4e6c53555139496e687463433570615751365155593552544a44526a6b" +
	"774e7a677a4d5446464f54677751304d354e4559794e6b55774d7a51774e" +
	"4549694948687463447044636d5668644739795647397662443069515752" +
	"76596d5567554768766447397a6147397749454e544e5342586157356b62" +
	"33647a496a3467504868746345314e4f6b526c636d6c325a575247636d39" +
	"7449484e30556d566d4f6d6c7563335268626d4e6c53555139496e687463" +
	"433570615751364d6a63784e6a593251555977526a424752546b784d5546" +
	"464f4467354d5451314d7a49784d444d334f44596949484e30556d566d4f" +
	"6d5276593356745a57353053555139496e68746343356b615751364d6a45" +
	"784e6a593251555977526a424752546b784d5546464f4467354d5451314d" +
	"7a49784d444d334f4459694c7a3467504339795a4759365247567a59334a" +
	"70634852706232342b49447776636d526d4f6c4a45526a3467504339344f" +
	"6e68746347316c6447452b4944772f654842685932746c6443426c626d51" +
	"39496e4969507a372f3767414f51575276596d55415a4d4141414141422f" +
	"397341684141554542415a45686b6e4678636e4d6959664a6a49754a6959" +
	"6d4a69342b4e5455314e54552b5245464251554642515552455245524552" +
	"455245524552455245524552455245524552455245524552455245524552" +
	"454152555a4753416349435959474359324a69416d4e6b51324b79733252" +
	"455245516a56435245524552455245524552455245524552455245524552" +
	"455245524552455245524552455245524552455245524552455245542f77" +
	"41415243414541415141444153494141684542417845422f385141686741" +
	"414177454241514541414141414141414141414141414149444151514642" +
	"674542415145424141414141414141414141414141414141514143417841" +
	"4141514d4441674d46425167414241634141414141415141524169457841" +
	"3045535557454563594569457758776b6147783063486838544a43556851" +
	"565970496a4d344b6930754a445577595241514542415145414167494441" +
	"41414141414141414141424552494349544642555a4643412f2f61414177" +
	"444151414345514d52414438412b4d5168436b374f6b776562444b663278" +
	"2b2f37467978494677363976306e4732467a2b6f6e33572b7138584a4134" +
	"356d42754352376c30396563382b61783575327868695933444f745a5943" +
	"4e5672474a584a305636664d634d335a78615134684e3165415935626f56" +
	"684b735437616a56514b37756c6b4d7354676d7a532f4b6545766f667655" +
	"586e4954354d5a78794d4a5549534a5a4345495567684346494951685343" +
	"4549556e314f492f78765143662f666b2b332f73587a5939362b68395a49" +
	"773942306654436e673877397376786b7641424d6a74376768706b714275" +
	"395966434854477059424a6b4e57466c4a4e61355a74466945736843454b" +
	"515168436b72684564336a566a6c78534c4564374c6c425a577a62694249" +
	"736559586278367a7a636b2f68697a355a6d78624b69785556307a2f3251" +
	"36356c6e2f535358342f507966503076697843554a5a4a576a38536a4830" +
	"75544a413547323478517a6c51666565516372733657455a3449695a6148" +
	"6d446632657a4c7539637847585534634944596d614c577658344e32686a" +
	"71726e34572f4b2f54342f4b78526877433872724f696e6b7a7938706a4b" +
	"58693266713775505948504a64505664666b6a6b45635553514455746673" +
	"2b763258683668507a426a7a346e65334d48543756313932575a50367566" +
	"6d5758663234734f447a444b456e45774b44375062384a524f692b6f3953" +
	"786a7a634f5168737842337477322f3841555342795a664c35473346724f" +
	"7550727a6b64504e307773785452445853626f375266633965446166616a" +
	"66526c7a6448666e682f4c786561503841636750482f696a2b3775312f46" +
	"655971787a47424f776e6833465356425168434567495168534345495567" +
	"6d78774f535168473869414f394b685366522f38413075554872546a4835" +
	"635559597833422f745869694a59795073556d58714d6d6249637551764f" +
	"56544a54457445593171346f4449726d756d4d69517956496f5168436745" +
	"4951704243465848694567354c556b6663704a674f71526848395248636e" +
	"6830346b576639724d4f4b504947356764447077544c4a2b45544e6c3330" +
	"4667704c706e30323171334946766232376e5757474942496b37646e3253" +
	"4b72626274556d48365871664a654d71776c656a7231596452484c6a4f4d" +
	"53453859424f32622b48736c516a5839773475437641517454316e777a66" +
	"4f75387a79534c346e326b32384a7165595a7a3342646e53395152486347" +
	"335141416c4d6954442f44454d41315458636547713859794a414430466c" +
	"6b695a467a556c585335656e3150585247375a497a6e4b386a3744336257" +
	"43387462474a6c51423159644c6c5036543330576658726673795a394949" +
	"58522f437a66742b495444302f4d513766454c4f78724b35554c72487032" +
	"556c714476566836545039306669725973727a6b4c307836536631544863" +
	"466f394a422f582f792f657271486d764c517657486f7a33794e2f772f65" +
	"74506f6830794139794f6f75613868433975486f63522b655a505946622b" +
	"6f3663573348744b756f654b2b6551766f763666436173663879582b6a78" +
	"635a653866525855584666506f58305139447766756c3778394670394277" +
	"48387335447462376c6452635638346866516e304846706b507757663047" +
	"4c5849666372714c69766e304c336a364669482f6b502b564e2f53395038" +
	"41756d2f64394664516331382b68652f2f41452f546a5758762b354a5030" +
	"62462b6d636832682f6f72714c6d764456384d706745676752464b383132" +
	"7a39486b50797a42504d4e39564f507033555164747248762b6164677975" +
	"66486b6b386a516b654a792f77436e67694f575753596941412f67316174" +
	"46305239503669424a4456532f77656f4574344163566f797469796f354d" +
	"6b6f6d494c65476f6f3271334c484a6a6964306e42384f7632716b75687a" +
	"797542334d456b2b6a7a794a6c4b4c6b38777255674d55705743714f6c6b" +
	"626b42644b3046576e456f644944636c644f4c7073636441547a57417033" +
	"5754496f78464162634562696c4254752b69476d78634b676b624a596c72" +
	"424f4355457a67364a674f535545706b46574c4d793052694f7854415469" +
	"4a51544e464145526f73327074716948354c51546f417445552b794c5564" +
	"2b78534a756b7433466274354c64764a424b355135546d434e7167523172" +
	"7074764a477771536272435658797973324649534a4356574d4668696f49" +
	"705372454a4345684a696c4d5a6343716c4b6b4a45454a5353716c4b5568" +
	"35414b64494171526974417754694a37466b594f72787838566b6b415478" +
	"6956574d514c4a683249614c4742546a47363048697470336f5459516962" +
	"30544144524c46797141447355544349437834716371584b53596b434751" +
	"6464472f6b6e456c7a676d5849725161375361714c704530323552465256" +
	"4f4f4151546d52354b686c454339567a76566b3361704b76537978494a42" +
	"4d48556d6772585a44495a43593645774646724a424446316e6c71724c4e" +
	"716b6c734378686f756a796962304344434d655a53484d59704346305359" +
	"71524153796759705446564c4a4374423430594656414153622b43614a34" +
	"705a5769465146524254774c6f615742566d33564b6b4b56315462754b79" +
	"5652454932425a416771304168704d594872564f4f6e46335a574256496c" +
	"7531476e4568302b6a6f2f693158513530436f4c4930343454304a4f724a" +
	"6f2b6e534e584646334433726161713272484965686e72494d744852532f" +
	"63753042444f6a546a6b4852635a6642502f446a72497271596f41347131" +
	"59357830754d463946547934436754686a32724448336f52664c6a6f7952" +
	"6f717a4d6c49354a524141516c4d65436f526f734b676d344643686e736d" +
	"3755724d6c466b434568436f5136513053796d516b4c4b71517857676d51" +
	"6b494371516b4b512b654631514b55536e4257324641797045734b4b6358" +
	"4e6c6149653679513869725167347173413270393341456f61576a414255" +
	"69542b43684355724d417268774f4b793072466c546342565233536c6f49" +
	"7034745a3337554665424d6c554b556541564b43367930634f624a674572" +
	"416969594d3346424f7457426a554a6f417154467571326c796d62334b54" +
	"4e727165516b57695371694a304b777961703979676c416b3369514f4365" +
	"504d4c446c54524149715649624968314d7847685670524458555358464b" +
	"4b42534572634530716b454a6d43306b6a424b597175326a724347554544" +
	"464b597664574b56686461434a6a775579437179706f70376b73766d4146" +
	"534b324261365978424c73793277324264564d5a58305344754363486263" +
	"643643724435716f47366a4a49374a42773771774469374c4c5234685668" +
	"51574b54475147425968576a4679345130614942756e68474f6957394c4b" +
	"6b61303134724c5268466b2b314b51593046452b47464b6e635167744467" +
	"4c5153657859664e4a305a6c614554426e7567744465394d477373497275" +
	"5a61416531314962646f76666d744a5035526445676543306e626f684e68" +
	"4c614b706a49714d70456c68336f42492b61677151365339714c4a5a4746" +
	"71636b6b5a476452546b556f474a55736f387547346d7872477475584e64" +
	"414f34564651744b51354d57555a5046473359516669716736706a414f36" +
	"47347151536b6f4b53526136514a636c4d706a524b536b456c52546c4631" +
	"51684b51744d766d575478484e4c74634d6d694b746f74734e4553545542" +
	"315552424166524b4a4777466445324f6c304e4b77614a6179723467374d" +
	"7067675363436959796f64466b7251494e5a4b3479415773755546714770" +
	"54527258514961644a6e4f52384e6c30695169474138584663555a4d4c36" +
	"7270693867794b31463479656a4256787841385430433534304c6d696f34" +
	"694f6179585649734b366f4a3442317a786d4352754e4b6c31574d714141" +
	"2b354255424c75433456444963464f4a5a6d466c736e436b6f4343454567" +
	"71635a474d5761694e394541675a336b4b6a6b6c4d433767305648733130" +
	"4766424b624137566b6a757271746255634b7053524837564a6a37617259" +
	"7939366c4c494843634d54785344733553454c646d723979576231556d45" +
	"715a544351535353474575464d6c6b756155322f77424a6a49614573764a" +
	"366a31664c674c5a634a6a7a33552b53334a724e7550596669736b426465" +
	"66673957775a76316254776c395632376e5667664e6d5a706f6d69534650" +
	"63316446534a417174734867474c75657856446e74555845533659546444" +
	"5330534766564f493772466c5048485943493856574b43654d535251747a" +
	"56576f7a5757517147494954514c304c686c6b6c686a65516b5255633131" +
	"516d314251446975636b51766234716b5337384371745258644b4e64447a" +
	"585444494a65473558435a426e657569734247544568794e65437958572b" +
	"32394f315842654c61385677695a317172516b59327167726d455978634f" +
	"364a78384f346c32344f6b47513349576562576c6967716b41732b69774d" +
	"6462364c504d4556753453464456515941596e694757376a594c42505368" +
	"574d446275556c474a7141706949334237706849783746676b4451715244" +
	"6945693969743274593254303071706d6a3755686e386830436645753677" +
	"694d695333616c4d41365171575a4959685955706e566b6f704458756b6d" +
	"42494d51344b636c366170536c6c355056656a59636c6366676c7973764a" +
	"6c4c71665435625849476e375376715358554d324b4f614f7a494843334b" +
	"785a2b6e7a316b305a4e32704b707751566f4b4f35357172454652455246" +
	"6c574a4242696464466b7178613456495345614571474f6172746a6b4730" +
	"31396e5130766a65526342786f6e4a7233714f4d67436c6971754b684250" +
	"416758484e6b3444467852317a7742446952634f38667839766d722b492b" +
	"493169654b435566366b76454e7046693636593767435a436d6a4b517951" +
	"795366583470786b656850626f67746a6b61725756595a42747578553344" +
	"41477a657a4a495241446b714c75696461464d42486a37317a5266696e47" +
	"586a775753615753416b496b3936704f4f3041784955795262354a74776c" +
	"72513055674464376a5662475a46456b67774c5066526246694842725253" +
	"4e4c4b7a6b304957526d43654253376e6574466747303878645157684a79" +
	"5449696c696745477864495258525a733231466b6f53754344566b623656" +
	"7155587173634f586f36517a6348575853354e7074324c597733436d6951" +
	"772f464b5841436a6c794d526f5673637646494d344b7767684d414a3254" +
	"4130597044354f4a32326f6e6a5435705763494c6738467468614c614a37" +
	"56754567445647716550695961494b6b432f4e4d2f44515639765971654d" +
	"73547a487956494861484b437144756b525a6c6f6d49796f6566324f6b2f" +
	"4c51364145707052456751474a41515851596b5978496e63434e42795242" +
	"774b466e31344c6c47536f693944386b2b504e356b694132375772494f75" +
	"316d414a76785443424c74554836704d65514d353461707a495641655063" +
	"687075556e486a4f615949786767536b3346626a6b537869784246466b38" +
	"307a45776c574653326a4a544a774e74767238304a595342346a526c7368" +
	"522b78536a4e366d347256556849786375344b69574c74577854695a4248" +
	"4170524a694f50377531614a4f58485a7935715373706b6b6b4e75436173" +
	"3647363569646a6b36703856727152695343784873564f514a4c2f704655" +
	"356b394465793078416938545671715459354a43494f76305654496b3046" +
	"57717062596755737359786b34754b4b523479456d30644a4e366e687753" +
	"7a6b42625652387754426944644c494751695242344a764e32446b705a54" +
	"54644c2f4d6b4d7a4734716b4c544979654b7734714d5352517243514377" +
	"73736b624f6b4f6d456a453153355a62532b677564464d5a4769774c6a6d" +
	"6e6a494d51624a443573794a4b714b6869334253596b55756e6744395670" +
	"6853494938514e717079444567684b434b67684d4274484d496144304a69" +
	"4c2b337a56537a4d4f314a413676645043514a414272564262756c614651" +
	"617571593553334f6539546261547a6f566f6b776356487848334b4b6759" +
	"6c6d626833716768476f31736c6850636550337132783668756179537732" +
	"7a63456b454832373151793231633857497270386c43556953346f415651" +
	"47385a42774c6356492b484e435141485a54322f4657324d6156424f7134" +
	"6a686470596d654e3331484c6e3272706a4d435a694859436c5546544d30" +
	"6153444774514e5541737834705833674756615835724a6e777337435465" +
	"3952554263474a464470783972496930584955686b4553597a4c56483370" +
	"70786c47704e395649354c754c364a72637450626b567a794c6c7a546d71" +
	"53796267356f51667a4b523478456753346f61386b38524b68346a696f48" +
	"4a744a6c594d79726a6b53484e5278556a753763446f5532376b34555153" +
	"443855776b6579587a556c4a4557584c6b4f50634137534a5a564d6c676b" +
	"414b2b39496332664673424149494e77706559344d433462526463346275" +
	"3563306f7651684d5a414c697431686369716d4159574e4854536d7a326f" +
	"6b4a2b5a4b446267726a50455233457346353356656f52413234716e6a6f" +
	"764d6e6b6c4d76497574597a727167474470374b494c4e32716a6a337142" +
	"344f4a564b7134335072645371506d7442646e5556774b4d6b6c50623468" +
	"56316d34697353375757694a7633494b78724639566a526b2f7a53526b57" +
	"5a2b777073636f6b316a596539425847485548784d6d4a6b61436c48584e" +
	"444b496b6a536972697a5642314874373046654d326c573946704f77474a" +
	"4f6f767a344a4a536f416267756539556e4c66346832565152754d504561" +
	"673054596a756948763841426b67446e67496b73654b7041445a334b4c59" +
	"795049456c766373794555424642566c7351436145734b704a544d532b70" +
	"6453615045387a5869654a4b6d444b443478523763505a6c534e47705856" +
	"6b6f7251466a5a53574f61457749476b6a516f78435559796937676e6238" +
	"474c71496935304a747a577649416d4e72653375556a5373326c69732f49" +
	"4b42673975584a615a2f75413851707a5275464259476755447a66744341" +
	"51576374514b557a355138566c6b4348414a426669704c79334131346355" +
	"7368756f6d6d5769787370487773316a386b6f305a474e45453155664e4d" +
	"433072636543686d395277776f44753745344e587a546a69675a534c4336" +
	"384c7165726c6e706150425a315856533669546d3267584f74794f646f51" +
	"68435758554530564f4b6363316c70523645706758706f70416b47697275" +
	"44714c5938685650435250636a5a2b704c516173676d6a4a2f7936617035" +
	"6b56497550745752327969445931634b637043353769704d6d6439654846" +
	"4a4358747853797147466973736548344a546f686d637356635a37636c77" +
	"3141636933327034687930533342423136486e7437635661505541454e61" +
	"6e65764a7835435152472f4432397171735a57356c474858716e4d493034" +
	"6969356a6c496b2b696c4c4951474f696c504a734c364769734e72734578" +
	"51646f597035457841656f5a31776a714d654a3930726142527a6571796c" +
	"5447473571776249394f65336275656c484b7945336c4b45785236536632" +
	"3058697736374e41754a58756d48714f5146797854794f6f397779327343" +
	"6158393655784249463947586b2f32556e6661416255736e2f7441627852" +
	"6c5855656f447569596c7944384548612f693046313573665534682b616f" +
	"50557354317372467364354c45682b594b77354e6b586b34414861755139" +
	"64696b4b53727a553833583434784f32704f69735775487175736c6e4c57" +
	"6a6f46797253584c72463063676843464949516853586a4b696345674263" +
	"346d79635a427a51566f79633153524a4536724e34564d65534a754b3856" +
	"46627a445274466f6b5a32734b704a4d5274534573514b75677569556949" +
	"6e69464a3369783157536d376764715277657a5253614e51736b64745233" +
	"646c575137462b5866565a4b5148356263314a5861306477725777753330" +
	"57776655505832397550776c4351684531754743542b5177414173706136" +
	"6874506a305762786a756671754f576152444367344253546731317a3676" +
	"396f39363535354a544c6c49684f44516843464149516853434549556768" +
	"4346494951685343454955676843464949516853662f5a"
