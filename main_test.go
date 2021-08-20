package main

import (
	"testing"
)

const operator_acreage_wkt = `MULTIPOLYGON (((-93.47619450000001 35.7018964, -93.47622610000001 35.7005955, -93.47631389999999 35.696985, -93.4764016 35.6933745, -93.47648940000001 35.6897639, -93.4808051 35.6898963, -93.4851208 35.6900287, -93.485045 35.6936293, -93.48496919999999 35.6972299, -93.4848935 35.7008305, -93.4848643 35.7022189, -93.47619450000001 35.7018964)), ((-93.4195691 35.6997902, -93.4195812 35.6988627, -93.4217534 35.6989225, -93.421744 35.6998711, -93.4195691 35.6997902)), ((-93.3450994 35.6970473, -93.3452046 35.6945325, -93.34520449999999 35.6945325, -93.3404381 35.6944917, -93.34059190000001 35.6908345, -93.3453581 35.6908631, -93.3497704 35.690919, -93.3496224 35.6945819, -93.3495173 35.6971846, -93.3465465 35.6970741, -93.3450994 35.6970473)), ((-93.6317102 35.6723863, -93.6316357 35.6741593, -93.6318729 35.6740104, -93.6324619 35.6737364, -93.6351009 35.6737834, -93.6357469 35.6737144, -93.63605010000001 35.673616, -93.6360567 35.673616, -93.6361071 35.6724774, -93.6317102 35.6723863, -93.6318598 35.6688284, -93.6362649 35.6689075, -93.64067009999999 35.6689866, -93.64507519999999 35.6690657, -93.6452496 35.6654718, -93.64959380000001 35.6655932, -93.6539381 35.6657146, -93.6539381 35.6657147, -93.6537957 35.6693033, -93.65815600000001 35.669422, -93.66251629999999 35.6695408, -93.66467280000001 35.6696142, -93.6645673 35.6731664, -93.6645637 35.6732077, -93.6668464 35.6732961, -93.6667334 35.6769002, -93.6622958 35.6767074, -93.6622958 35.6767074, -93.6579035 35.6765939, -93.6577772 35.6801799, -93.6533688 35.6800691, -93.64896040000001 35.6799584, -93.6491187 35.6763671, -93.64927710000001 35.6727758, -93.6449008 35.6726597, -93.6405039 35.6725686, -93.6403378 35.6761505, -93.6359492 35.6760473, -93.6357914 35.6796172, -93.63141109999999 35.6795019, -93.63132349999999 35.6830709, -93.6357213 35.6831809, -93.63565130000001 35.6867447, -93.6355812 35.6903085, -93.6355112 35.6938723, -93.6310608 35.693778, -93.6310608 35.6937778, -93.63114830000001 35.6902089, -93.62671539999999 35.6901094, -93.6268206 35.6865351, -93.6269257 35.6829609, -93.62254470000001 35.6828529, -93.622647 35.6792713, -93.61826309999999 35.679156, -93.61387929999999 35.6790406, -93.61402099999999 35.6754727, -93.6184047 35.6755954, -93.6227884 35.6757182, -93.6271721 35.6758409, -93.62731340000001 35.6722951, -93.6317102 35.6723863), (-93.650976 35.6697243, -93.6494138 35.6696763, -93.64935180000001 35.6710826, -93.650927 35.6711139, -93.650976 35.6697243)), ((-93.6004703 35.6895156, -93.6005751 35.6859118, -93.6049417 35.6860222, -93.60504090000001 35.6824185, -93.60940189999999 35.682529, -93.6093083 35.6861327, -93.6092147 35.6897364, -93.6091211 35.6933401, -93.6047433 35.6932297, -93.6048425 35.689626, -93.6004703 35.6895156)), ((-93.52993789999999 35.6840534, -93.5301001 35.6804169, -93.534621 35.6805461, -93.5344561 35.6841754, -93.5342912 35.6878048, -93.5297758 35.6876899, -93.52993789999999 35.6840534)), ((-93.47241440000001 35.6860246, -93.4726551 35.6824176, -93.4769395 35.68256, -93.4771645 35.678958, -93.4773895 35.6753561, -93.48164250000001 35.6755085, -93.4858954 35.675661, -93.48570170000001 35.6792529, -93.4855081 35.6828448, -93.485508 35.6828448, -93.4812238 35.6827024, -93.48101440000001 35.6862994, -93.47671440000001 35.686162, -93.47241440000001 35.6860246)), ((-93.50345059999999 35.6760352, -93.50789159999999 35.6761533, -93.5123327 35.6762715, -93.5123327 35.6762716, -93.51226629999999 35.6780116, -93.5127598 35.6782724, -93.5135169 35.6790974, -93.5139659 35.6794645, -93.51438690000001 35.6796245, -93.5166299 35.6799402, -93.51662589999999 35.6800379, -93.5164781 35.6836863, -93.5120546 35.6835631, -93.50763120000001 35.68344, -93.50776140000001 35.6797967, -93.5033291 35.6796761, -93.50345059999999 35.6760352)), ((-93.5915996 35.6768951, -93.5921414 35.6766353, -93.5920739 35.6784832, -93.5894019 35.6784359, -93.5897659 35.6778789, -93.5901977 35.6773281, -93.59055499999999 35.677209, -93.5912994 35.6770452, -93.5915996 35.6768951)), ((-93.50345059999999 35.6760352, -93.49906180000001 35.6759416, -93.49467300000001 35.6758481, -93.4948054 35.672229, -93.4904042 35.6721369, -93.4860029 35.6720449, -93.48600879999999 35.6718484, -93.4865518 35.6717455, -93.4876179 35.6712185, -93.4888809 35.6708285, -93.4895818 35.6704844, -93.4904619 35.670397, -93.4905088 35.6703924, -93.4909579 35.6704614, -93.4928098 35.6704374, -93.4946059 35.6706424, -93.49486020000001 35.6707311, -93.4949378 35.6686098, -93.4993515 35.6687004, -93.49920659999999 35.672321, -93.50360790000001 35.672413, -93.50345059999999 35.6760352)), ((-93.5037651 35.6687909, -93.50392239999999 35.6651687, -93.50836289999999 35.6652942, -93.5128033 35.6654196, -93.5126465 35.6690369, -93.5170871 35.6691599, -93.51693040000001 35.6727748, -93.51248959999999 35.6726542, -93.5080487 35.6725336, -93.503608 35.6724103, -93.5037651 35.6687909)), ((-93.46443120000001 35.6711752, -93.4645044 35.6676066, -93.46898710000001 35.6678699, -93.46893540000001 35.6714605, -93.46443120000001 35.6711752)), ((-93.6143045 35.6683368, -93.6144462 35.6647689, -93.6188295 35.6649138, -93.6232127 35.6650587, -93.62307130000001 35.6686118, -93.6186879 35.6684743, -93.6143045 35.6683368)), ((-93.42454290000001 35.6662634, -93.4246325 35.6626525, -93.42906429999999 35.6627305, -93.4289567 35.6663364, -93.42454290000001 35.6662634)), ((-93.3721915 35.6656328, -93.37228880000001 35.6619843, -93.3723861 35.6583358, -93.37668669999999 35.658368, -93.3790829 35.6583859, -93.3789929 35.6620377, -93.3765867 35.6620185, -93.37648660000001 35.6656691, -93.3721915 35.6656328)), ((-93.49549279999999 35.6541527, -93.4956378 35.6505432, -93.5000769 35.6506366, -93.4999316 35.6542441, -93.50437049999999 35.6543355, -93.50437049999999 35.6543355, -93.5042251 35.657941, -93.5040796 35.6615465, -93.49964110000001 35.6614591, -93.4997864 35.6578516, -93.4953477 35.6577621, -93.49549279999999 35.6541527)), ((-93.48675969999999 35.6503564, -93.4869044 35.646743, -93.4913437 35.6468384, -93.49119880000001 35.6504498, -93.4910539 35.6540613, -93.486615 35.6539698, -93.48675969999999 35.6503564)), ((-93.45597239999999 35.6491505, -93.4560395 35.6455258, -93.46011489999999 35.6457339, -93.4605462 35.6457559, -93.4644059 35.6459529, -93.4650528 35.645986, -93.4649524 35.6496018, -93.4648519 35.6532243, -93.4603786 35.6530009, -93.4604624 35.6493761, -93.45597239999999 35.6491505)), ((-93.69107889999999 35.6488962, -93.6911868 35.6453499, -93.69569920000001 35.6454688, -93.69559940000001 35.6490223, -93.69107889999999 35.6488962)), ((-93.5002222 35.6470292, -93.5003111 35.6434455, -93.5047349 35.6435457, -93.5046614 35.6471246, -93.5002222 35.6470292)), ((-93.42500769999999 35.6446235, -93.42509769999999 35.64109, -93.4295396 35.641196, -93.4339815 35.641302, -93.43391130000001 35.6448441, -93.42945949999999 35.6447338, -93.42500769999999 35.6446235)), ((-93.4915835 35.6396517, -93.4914636 35.643245, -93.48703980000001 35.6431448, -93.4825142 35.6430047, -93.4826759 35.6394092, -93.4828376 35.6358136, -93.48731050000001 35.6359485, -93.4871751 35.6395466, -93.4915835 35.6396517)), ((-93.464612 35.6423714, -93.4648181 35.63879, -93.4691783 35.6389968, -93.4691783 35.6389968, -93.4689376 35.6425844, -93.464612 35.6423714)), ((-93.3328697 35.6402313, -93.3283916 35.6402207, -93.3283916 35.6402207, -93.3328659 35.636779, -93.3328659 35.6367781, -93.33286699999999 35.6367781, -93.3373369 35.6333397, -93.3373423 35.6367908, -93.3373477 35.6402418, -93.3328697 35.6402313)), ((-93.4915835 35.6396517, -93.49170340000001 35.6360583, -93.4960962 35.6361682, -93.4959918 35.6397567, -93.4915835 35.6396517)), ((-93.3008945 35.6331776, -93.30090010000001 35.6296618, -93.3056026 35.6297484, -93.31030509999999 35.6298351, -93.3102411 35.633362, -93.310177 35.636889, -93.305533 35.6367912, -93.30556780000001 35.6332698, -93.3008945 35.6331776)), ((-93.28717090000001 35.6328794, -93.2871157 35.6293513, -93.291495 35.6294884, -93.29154800000001 35.6329932, -93.29154800000001 35.6329932, -93.28717090000001 35.6328794)), ((-93.3373646 35.626257, -93.3373977 35.6226253, -93.3418827 35.6226513, -93.34635969999999 35.6226782, -93.3463677 35.6226773, -93.3463856 35.6212372, -93.3477259 35.6227032, -93.3477955 35.6227793, -93.34778900000001 35.6245758, -93.3463837 35.6245854, -93.34634029999999 35.6245838, -93.3463228 35.6263, -93.3462779 35.6299227, -93.3418047 35.6299057, -93.3373315 35.6298887, -93.3373646 35.626257)), ((-93.4431922 35.6273265, -93.4433174 35.6237937, -93.44776659999999 35.6238954, -93.4476356 35.6274233, -93.4431922 35.6273265)), ((-93.4876586 35.6251208, -93.4877649 35.6215061, -93.4878712 35.6178914, -93.4922634 35.6179604, -93.4966556 35.6180295, -93.4965419 35.621667, -93.49215340000001 35.6215866, -93.49204330000001 35.6252127, -93.4876586 35.6251208)), ((-93.44341249999999 35.6211111, -93.4434426 35.6202609, -93.4478977 35.6203675, -93.447867 35.6211933, -93.44341249999999 35.6211111)), ((-93.6833668 35.6129055, -93.68351730000001 35.6092918, -93.6880007 35.6094232, -93.6878476 35.6130321, -93.6923283 35.6131587, -93.69217260000001 35.6167629, -93.6876944 35.6166411, -93.6832163 35.6165193, -93.6833668 35.6129055)), ((-93.43931449999999 35.6093117, -93.4437569 35.6094183, -93.4436623 35.6130732, -93.4435678 35.616728, -93.4435678 35.6167281, -93.4391068 35.6166166, -93.43464590000001 35.6165051, -93.434759 35.612855, -93.4392107 35.6129641, -93.43931449999999 35.6093117)), ((-93.3795831 35.616535, -93.3797732 35.6128505, -93.3819747 35.6129404, -93.3818484 35.6158603, -93.38185129999999 35.6166433, -93.3795831 35.616535)), ((-93.3731246 35.6125791, -93.37329459999999 35.6089395, -93.37772390000001 35.6090896, -93.3775634 35.6127603, -93.37740290000001 35.616431, -93.37511910000001 35.616322, -93.3753014 35.612668, -93.3731246 35.6125791)), ((-93.5321244 35.6116023, -93.5322625 35.6079757, -93.5367376 35.6080958, -93.54121259999999 35.6082158, -93.54105680000001 35.611847, -93.54090100000001 35.6154781, -93.5364436 35.6153535, -93.5319863 35.6152289, -93.5321244 35.6116023)), ((-93.29720829999999 35.6076585, -93.29733640000001 35.6039954, -93.3019078 35.604075, -93.3064792 35.6041546, -93.3063334 35.607803, -93.3061876 35.6114514, -93.3016339 35.6113865, -93.2970802 35.6113216, -93.29720829999999 35.6076585)), ((-93.43931449999999 35.6093117, -93.43487210000001 35.609205, -93.43498510000001 35.605555, -93.4350982 35.601905, -93.43519980000001 35.5984018, -93.4396295 35.5984913, -93.4440591 35.5985807, -93.443946 35.6021086, -93.4395221 35.6020068, -93.4394183 35.6056592, -93.43931449999999 35.6093117)), ((-93.27494040000001 35.601891, -93.27497889999999 35.6002411, -93.2766474 35.6002415, -93.2766398 35.6018847, -93.27494040000001 35.601891)), ((-93.63936870000001 35.6004046, -93.6394488 35.596825, -93.6395289 35.5932455, -93.6439673 35.5933892, -93.6484057 35.5935329, -93.64832579999999 35.5970981, -93.6438873 35.5969615, -93.64380730000001 35.6005339, -93.63936870000001 35.6004046)), ((-93.4218365 35.594667, -93.4219837 35.5911127, -93.4264567 35.5912069, -93.4263248 35.5947442, -93.4218365 35.594667)), ((-93.5154222 35.5858819, -93.51556100000001 35.582294, -93.5156999 35.5787061, -93.5201325 35.5788217, -93.5245652 35.5789374, -93.52439680000001 35.5825401, -93.52439680000001 35.5825401, -93.5199789 35.582417, -93.51982529999999 35.5860123, -93.5196718 35.5896076, -93.51528329999999 35.5894698, -93.5154222 35.5858819)), ((-93.3738603 35.5797555, -93.37394620000001 35.5761096, -93.3783948 35.5761681, -93.37832400000001 35.5798129, -93.3738603 35.5797555)), ((-93.6445102 35.5753352, -93.6443591 35.5789926, -93.6398974 35.5788679, -93.6400501 35.5752096, -93.6445102 35.5753352)), ((-93.67575530000001 35.5762856, -93.6758873 35.5726167, -93.6714065 35.5724793, -93.6715358 35.5688178, -93.6716658 35.5651351, -93.67256930000001 35.5651612, -93.6761518 35.5652648, -93.67601929999999 35.5689478, -93.6805028 35.5690778, -93.6849862 35.5692077, -93.68498630000001 35.5692077, -93.6848488 35.5728915, -93.6803681 35.5727541, -93.6802333 35.5764304, -93.67575530000001 35.5762856)), ((-93.6445102 35.5753352, -93.6446613 35.5716777, -93.64481240000001 35.5680203, -93.6492693 35.5681476, -93.64911979999999 35.5718042, -93.6489703 35.5754607, -93.64897019999999 35.5754607, -93.6445102 35.5753352)), ((-93.4097536 35.5656577, -93.4098436 35.5620206, -93.4142762 35.5620927, -93.4143859 35.5584617, -93.4188381 35.5585398, -93.4187088 35.5621648, -93.4185796 35.5657898, -93.4184504 35.5694148, -93.41832119999999 35.5730398, -93.4139473 35.5729858, -93.4139473 35.5729858, -93.414057 35.5693548, -93.40966349999999 35.5692947, -93.4097536 35.5656577)), ((-93.6640828 35.5505284, -93.6642093 35.5468595, -93.66875330000001 35.5469774, -93.66875469999999 35.5469774, -93.6686526 35.5506511, -93.6684891 35.5542456, -93.6683256 35.5578402, -93.6683256 35.5578401, -93.66371820000001 35.5577115, -93.6639005 35.5541199, -93.6640828 35.5505284)), ((-93.54220429999999 35.5540032, -93.5424076 35.5503807, -93.54678490000001 35.5505196, -93.54657280000001 35.5541444, -93.54636069999999 35.5577692, -93.542001 35.5576257, -93.54220429999999 35.5540032)), ((-93.22224730000001 35.5338452, -93.2223106 35.5302449, -93.22674120000001 35.5303486, -93.2266781 35.5339614, -93.226615 35.5375741, -93.2221839 35.5374454, -93.22224730000001 35.5338452)), ((-93.57921949999999 35.522365, -93.5794092 35.5187139, -93.5841059 35.5189106, -93.5841059 35.5189107, -93.5839237 35.5225483, -93.5886278 35.5227316, -93.58845289999999 35.526356, -93.58827789999999 35.5299803, -93.5835591 35.5298238, -93.58218890000001 35.5297783, -93.582407 35.5261379, -93.57902989999999 35.5260161, -93.57921949999999 35.522365)), ((-93.2278891 35.4795276, -93.22797540000001 35.4759155, -93.2324026 35.476022, -93.2368298 35.4761285, -93.2367423 35.4797636, -93.2323157 35.4796456, -93.2278891 35.4795276)), ((-93.4979824 35.4238753, -93.500011 35.423826, -93.5000209 35.4255059, -93.4980042 35.4254679, -93.4979824 35.4238753)), ((-93.3156175 35.3792996, -93.3157301 35.3756666, -93.3202034 35.375679, -93.3246766 35.3756914, -93.3245581 35.3793203, -93.32443960000001 35.3829492, -93.3199722 35.3829409, -93.3200878 35.37931, -93.3156175 35.3792996)))`

const operator_acreage_geojson = `{"type":"Feature","geometry":{"type":"MultiPolygon","coordinates":[[[[-93.4761945,35.7018964],[-93.4762261,35.7005955],[-93.4763139,35.696985],[-93.4764016,35.6933745],[-93.4764894,35.6897639],[-93.4808051,35.6898963],[-93.4851208,35.6900287],[-93.485045,35.6936293],[-93.4849692,35.6972299],[-93.4848935,35.7008305],[-93.4848643,35.7022189],[-93.4761945,35.7018964]]],[[[-93.4195691,35.6997902],[-93.4195812,35.6988627],[-93.4217534,35.6989225],[-93.421744,35.6998711],[-93.4195691,35.6997902]]],[[[-93.3450994,35.6970473],[-93.3452046,35.6945325],[-93.3452045,35.6945325],[-93.3404381,35.6944917],[-93.3405919,35.6908345],[-93.3453581,35.6908631],[-93.3497704,35.690919],[-93.3496224,35.6945819],[-93.3495173,35.6971846],[-93.3465465,35.6970741],[-93.3450994,35.6970473]]],[[[-93.6317102,35.6723863],[-93.6316357,35.6741593],[-93.6318729,35.6740104],[-93.6324619,35.6737364],[-93.6351009,35.6737834],[-93.6357469,35.6737144],[-93.6360501,35.673616],[-93.6360567,35.673616],[-93.6361071,35.6724774],[-93.6317102,35.6723863],[-93.6318598,35.6688284],[-93.6362649,35.6689075],[-93.6406701,35.6689866],[-93.6450752,35.6690657],[-93.6452496,35.6654718],[-93.6495938,35.6655932],[-93.6539381,35.6657146],[-93.6539381,35.6657147],[-93.6537957,35.6693033],[-93.658156,35.669422],[-93.6625163,35.6695408],[-93.6646728,35.6696142],[-93.6645673,35.6731664],[-93.6645637,35.6732077],[-93.6668464,35.6732961],[-93.6667334,35.6769002],[-93.6622958,35.6767074],[-93.6622958,35.6767074],[-93.6579035,35.6765939],[-93.6577772,35.6801799],[-93.6533688,35.6800691],[-93.6489604,35.6799584],[-93.6491187,35.6763671],[-93.6492771,35.6727758],[-93.6449008,35.6726597],[-93.6405039,35.6725686],[-93.6403378,35.6761505],[-93.6359492,35.6760473],[-93.6357914,35.6796172],[-93.6314111,35.6795019],[-93.6313235,35.6830709],[-93.6357213,35.6831809],[-93.6356513,35.6867447],[-93.6355812,35.6903085],[-93.6355112,35.6938723],[-93.6310608,35.693778],[-93.6310608,35.6937778],[-93.6311483,35.6902089],[-93.6267154,35.6901094],[-93.6268206,35.6865351],[-93.6269257,35.6829609],[-93.6225447,35.6828529],[-93.622647,35.6792713],[-93.6182631,35.679156],[-93.6138793,35.6790406],[-93.614021,35.6754727],[-93.6184047,35.6755954],[-93.6227884,35.6757182],[-93.6271721,35.6758409],[-93.6273134,35.6722951],[-93.6317102,35.6723863]],[[-93.650976,35.6697243],[-93.6494138,35.6696763],[-93.6493518,35.6710826],[-93.650927,35.6711139],[-93.650976,35.6697243]]],[[[-93.6004703,35.6895156],[-93.6005751,35.6859118],[-93.6049417,35.6860222],[-93.6050409,35.6824185],[-93.6094019,35.682529],[-93.6093083,35.6861327],[-93.6092147,35.6897364],[-93.6091211,35.6933401],[-93.6047433,35.6932297],[-93.6048425,35.689626],[-93.6004703,35.6895156]]],[[[-93.5299379,35.6840534],[-93.5301001,35.6804169],[-93.534621,35.6805461],[-93.5344561,35.6841754],[-93.5342912,35.6878048],[-93.5297758,35.6876899],[-93.5299379,35.6840534]]],[[[-93.4724144,35.6860246],[-93.4726551,35.6824176],[-93.4769395,35.68256],[-93.4771645,35.678958],[-93.4773895,35.6753561],[-93.4816425,35.6755085],[-93.4858954,35.675661],[-93.4857017,35.6792529],[-93.4855081,35.6828448],[-93.485508,35.6828448],[-93.4812238,35.6827024],[-93.4810144,35.6862994],[-93.4767144,35.686162],[-93.4724144,35.6860246]]],[[[-93.5034506,35.6760352],[-93.5078916,35.6761533],[-93.5123327,35.6762715],[-93.5123327,35.6762716],[-93.5122663,35.6780116],[-93.5127598,35.6782724],[-93.5135169,35.6790974],[-93.5139659,35.6794645],[-93.5143869,35.6796245],[-93.5166299,35.6799402],[-93.5166259,35.6800379],[-93.5164781,35.6836863],[-93.5120546,35.6835631],[-93.5076312,35.68344],[-93.5077614,35.6797967],[-93.5033291,35.6796761],[-93.5034506,35.6760352]]],[[[-93.5915996,35.6768951],[-93.5921414,35.6766353],[-93.5920739,35.6784832],[-93.5894019,35.6784359],[-93.5897659,35.6778789],[-93.5901977,35.6773281],[-93.590555,35.677209],[-93.5912994,35.6770452],[-93.5915996,35.6768951]]],[[[-93.5034506,35.6760352],[-93.4990618,35.6759416],[-93.494673,35.6758481],[-93.4948054,35.672229],[-93.4904042,35.6721369],[-93.4860029,35.6720449],[-93.4860088,35.6718484],[-93.4865518,35.6717455],[-93.4876179,35.6712185],[-93.4888809,35.6708285],[-93.4895818,35.6704844],[-93.4904619,35.670397],[-93.4905088,35.6703924],[-93.4909579,35.6704614],[-93.4928098,35.6704374],[-93.4946059,35.6706424],[-93.4948602,35.6707311],[-93.4949378,35.6686098],[-93.4993515,35.6687004],[-93.4992066,35.672321],[-93.5036079,35.672413],[-93.5034506,35.6760352]]],[[[-93.5037651,35.6687909],[-93.5039224,35.6651687],[-93.5083629,35.6652942],[-93.5128033,35.6654196],[-93.5126465,35.6690369],[-93.5170871,35.6691599],[-93.5169304,35.6727748],[-93.5124896,35.6726542],[-93.5080487,35.6725336],[-93.503608,35.6724103],[-93.5037651,35.6687909]]],[[[-93.4644312,35.6711752],[-93.4645044,35.6676066],[-93.4689871,35.6678699],[-93.4689354,35.6714605],[-93.4644312,35.6711752]]],[[[-93.6143045,35.6683368],[-93.6144462,35.6647689],[-93.6188295,35.6649138],[-93.6232127,35.6650587],[-93.6230713,35.6686118],[-93.6186879,35.6684743],[-93.6143045,35.6683368]]],[[[-93.4245429,35.6662634],[-93.4246325,35.6626525],[-93.4290643,35.6627305],[-93.4289567,35.6663364],[-93.4245429,35.6662634]]],[[[-93.3721915,35.6656328],[-93.3722888,35.6619843],[-93.3723861,35.6583358],[-93.3766867,35.658368],[-93.3790829,35.6583859],[-93.3789929,35.6620377],[-93.3765867,35.6620185],[-93.3764866,35.6656691],[-93.3721915,35.6656328]]],[[[-93.4954928,35.6541527],[-93.4956378,35.6505432],[-93.5000769,35.6506366],[-93.4999316,35.6542441],[-93.5043705,35.6543355],[-93.5043705,35.6543355],[-93.5042251,35.657941],[-93.5040796,35.6615465],[-93.4996411,35.6614591],[-93.4997864,35.6578516],[-93.4953477,35.6577621],[-93.4954928,35.6541527]]],[[[-93.4867597,35.6503564],[-93.4869044,35.646743],[-93.4913437,35.6468384],[-93.4911988,35.6504498],[-93.4910539,35.6540613],[-93.486615,35.6539698],[-93.4867597,35.6503564]]],[[[-93.4559724,35.6491505],[-93.4560395,35.6455258],[-93.4601149,35.6457339],[-93.4605462,35.6457559],[-93.4644059,35.6459529],[-93.4650528,35.645986],[-93.4649524,35.6496018],[-93.4648519,35.6532243],[-93.4603786,35.6530009],[-93.4604624,35.6493761],[-93.4559724,35.6491505]]],[[[-93.6910789,35.6488962],[-93.6911868,35.6453499],[-93.6956992,35.6454688],[-93.6955994,35.6490223],[-93.6910789,35.6488962]]],[[[-93.5002222,35.6470292],[-93.5003111,35.6434455],[-93.5047349,35.6435457],[-93.5046614,35.6471246],[-93.5002222,35.6470292]]],[[[-93.4250077,35.6446235],[-93.4250977,35.64109],[-93.4295396,35.641196],[-93.4339815,35.641302],[-93.4339113,35.6448441],[-93.4294595,35.6447338],[-93.4250077,35.6446235]]],[[[-93.4915835,35.6396517],[-93.4914636,35.643245],[-93.4870398,35.6431448],[-93.4825142,35.6430047],[-93.4826759,35.6394092],[-93.4828376,35.6358136],[-93.4873105,35.6359485],[-93.4871751,35.6395466],[-93.4915835,35.6396517]]],[[[-93.464612,35.6423714],[-93.4648181,35.63879],[-93.4691783,35.6389968],[-93.4691783,35.6389968],[-93.4689376,35.6425844],[-93.464612,35.6423714]]],[[[-93.3328697,35.6402313],[-93.3283916,35.6402207],[-93.3283916,35.6402207],[-93.3328659,35.636779],[-93.3328659,35.6367781],[-93.332867,35.6367781],[-93.3373369,35.6333397],[-93.3373423,35.6367908],[-93.3373477,35.6402418],[-93.3328697,35.6402313]]],[[[-93.4915835,35.6396517],[-93.4917034,35.6360583],[-93.4960962,35.6361682],[-93.4959918,35.6397567],[-93.4915835,35.6396517]]],[[[-93.3008945,35.6331776],[-93.3009001,35.6296618],[-93.3056026,35.6297484],[-93.3103051,35.6298351],[-93.3102411,35.633362],[-93.310177,35.636889],[-93.305533,35.6367912],[-93.3055678,35.6332698],[-93.3008945,35.6331776]]],[[[-93.2871709,35.6328794],[-93.2871157,35.6293513],[-93.291495,35.6294884],[-93.291548,35.6329932],[-93.291548,35.6329932],[-93.2871709,35.6328794]]],[[[-93.3373646,35.626257],[-93.3373977,35.6226253],[-93.3418827,35.6226513],[-93.3463597,35.6226782],[-93.3463677,35.6226773],[-93.3463856,35.6212372],[-93.3477259,35.6227032],[-93.3477955,35.6227793],[-93.347789,35.6245758],[-93.3463837,35.6245854],[-93.3463403,35.6245838],[-93.3463228,35.6263],[-93.3462779,35.6299227],[-93.3418047,35.6299057],[-93.3373315,35.6298887],[-93.3373646,35.626257]]],[[[-93.4431922,35.6273265],[-93.4433174,35.6237937],[-93.4477666,35.6238954],[-93.4476356,35.6274233],[-93.4431922,35.6273265]]],[[[-93.4876586,35.6251208],[-93.4877649,35.6215061],[-93.4878712,35.6178914],[-93.4922634,35.6179604],[-93.4966556,35.6180295],[-93.4965419,35.621667],[-93.4921534,35.6215866],[-93.4920433,35.6252127],[-93.4876586,35.6251208]]],[[[-93.4434125,35.6211111],[-93.4434426,35.6202609],[-93.4478977,35.6203675],[-93.447867,35.6211933],[-93.4434125,35.6211111]]],[[[-93.6833668,35.6129055],[-93.6835173,35.6092918],[-93.6880007,35.6094232],[-93.6878476,35.6130321],[-93.6923283,35.6131587],[-93.6921726,35.6167629],[-93.6876944,35.6166411],[-93.6832163,35.6165193],[-93.6833668,35.6129055]]],[[[-93.4393145,35.6093117],[-93.4437569,35.6094183],[-93.4436623,35.6130732],[-93.4435678,35.616728],[-93.4435678,35.6167281],[-93.4391068,35.6166166],[-93.4346459,35.6165051],[-93.434759,35.612855],[-93.4392107,35.6129641],[-93.4393145,35.6093117]]],[[[-93.3795831,35.616535],[-93.3797732,35.6128505],[-93.3819747,35.6129404],[-93.3818484,35.6158603],[-93.3818513,35.6166433],[-93.3795831,35.616535]]],[[[-93.3731246,35.6125791],[-93.3732946,35.6089395],[-93.3777239,35.6090896],[-93.3775634,35.6127603],[-93.3774029,35.616431],[-93.3751191,35.616322],[-93.3753014,35.612668],[-93.3731246,35.6125791]]],[[[-93.5321244,35.6116023],[-93.5322625,35.6079757],[-93.5367376,35.6080958],[-93.5412126,35.6082158],[-93.5410568,35.611847],[-93.540901,35.6154781],[-93.5364436,35.6153535],[-93.5319863,35.6152289],[-93.5321244,35.6116023]]],[[[-93.2972083,35.6076585],[-93.2973364,35.6039954],[-93.3019078,35.604075],[-93.3064792,35.6041546],[-93.3063334,35.607803],[-93.3061876,35.6114514],[-93.3016339,35.6113865],[-93.2970802,35.6113216],[-93.2972083,35.6076585]]],[[[-93.4393145,35.6093117],[-93.4348721,35.609205],[-93.4349851,35.605555],[-93.4350982,35.601905],[-93.4351998,35.5984018],[-93.4396295,35.5984913],[-93.4440591,35.5985807],[-93.443946,35.6021086],[-93.4395221,35.6020068],[-93.4394183,35.6056592],[-93.4393145,35.6093117]]],[[[-93.2749404,35.601891],[-93.2749789,35.6002411],[-93.2766474,35.6002415],[-93.2766398,35.6018847],[-93.2749404,35.601891]]],[[[-93.6393687,35.6004046],[-93.6394488,35.596825],[-93.6395289,35.5932455],[-93.6439673,35.5933892],[-93.6484057,35.5935329],[-93.6483258,35.5970981],[-93.6438873,35.5969615],[-93.6438073,35.6005339],[-93.6393687,35.6004046]]],[[[-93.4218365,35.594667],[-93.4219837,35.5911127],[-93.4264567,35.5912069],[-93.4263248,35.5947442],[-93.4218365,35.594667]]],[[[-93.5154222,35.5858819],[-93.515561,35.582294],[-93.5156999,35.5787061],[-93.5201325,35.5788217],[-93.5245652,35.5789374],[-93.5243968,35.5825401],[-93.5243968,35.5825401],[-93.5199789,35.582417],[-93.5198253,35.5860123],[-93.5196718,35.5896076],[-93.5152833,35.5894698],[-93.5154222,35.5858819]]],[[[-93.3738603,35.5797555],[-93.3739462,35.5761096],[-93.3783948,35.5761681],[-93.378324,35.5798129],[-93.3738603,35.5797555]]],[[[-93.6445102,35.5753352],[-93.6443591,35.5789926],[-93.6398974,35.5788679],[-93.6400501,35.5752096],[-93.6445102,35.5753352]]],[[[-93.6757553,35.5762856],[-93.6758873,35.5726167],[-93.6714065,35.5724793],[-93.6715358,35.5688178],[-93.6716658,35.5651351],[-93.6725693,35.5651612],[-93.6761518,35.5652648],[-93.6760193,35.5689478],[-93.6805028,35.5690778],[-93.6849862,35.5692077],[-93.6849863,35.5692077],[-93.6848488,35.5728915],[-93.6803681,35.5727541],[-93.6802333,35.5764304],[-93.6757553,35.5762856]]],[[[-93.6445102,35.5753352],[-93.6446613,35.5716777],[-93.6448124,35.5680203],[-93.6492693,35.5681476],[-93.6491198,35.5718042],[-93.6489703,35.5754607],[-93.6489702,35.5754607],[-93.6445102,35.5753352]]],[[[-93.4097536,35.5656577],[-93.4098436,35.5620206],[-93.4142762,35.5620927],[-93.4143859,35.5584617],[-93.4188381,35.5585398],[-93.4187088,35.5621648],[-93.4185796,35.5657898],[-93.4184504,35.5694148],[-93.4183212,35.5730398],[-93.4139473,35.5729858],[-93.4139473,35.5729858],[-93.414057,35.5693548],[-93.4096635,35.5692947],[-93.4097536,35.5656577]]],[[[-93.6640828,35.5505284],[-93.6642093,35.5468595],[-93.6687533,35.5469774],[-93.6687547,35.5469774],[-93.6686526,35.5506511],[-93.6684891,35.5542456],[-93.6683256,35.5578402],[-93.6683256,35.5578401],[-93.6637182,35.5577115],[-93.6639005,35.5541199],[-93.6640828,35.5505284]]],[[[-93.5422043,35.5540032],[-93.5424076,35.5503807],[-93.5467849,35.5505196],[-93.5465728,35.5541444],[-93.5463607,35.5577692],[-93.542001,35.5576257],[-93.5422043,35.5540032]]],[[[-93.2222473,35.5338452],[-93.2223106,35.5302449],[-93.2267412,35.5303486],[-93.2266781,35.5339614],[-93.226615,35.5375741],[-93.2221839,35.5374454],[-93.2222473,35.5338452]]],[[[-93.5792195,35.522365],[-93.5794092,35.5187139],[-93.5841059,35.5189106],[-93.5841059,35.5189107],[-93.5839237,35.5225483],[-93.5886278,35.5227316],[-93.5884529,35.526356],[-93.5882779,35.5299803],[-93.5835591,35.5298238],[-93.5821889,35.5297783],[-93.582407,35.5261379],[-93.5790299,35.5260161],[-93.5792195,35.522365]]],[[[-93.2278891,35.4795276],[-93.2279754,35.4759155],[-93.2324026,35.476022],[-93.2368298,35.4761285],[-93.2367423,35.4797636],[-93.2323157,35.4796456],[-93.2278891,35.4795276]]],[[[-93.4979824,35.4238753],[-93.500011,35.423826],[-93.5000209,35.4255059],[-93.4980042,35.4254679],[-93.4979824,35.4238753]]],[[[-93.3156175,35.3792996],[-93.3157301,35.3756666],[-93.3202034,35.375679],[-93.3246766,35.3756914],[-93.3245581,35.3793203],[-93.3244396,35.3829492],[-93.3199722,35.3829409],[-93.3200878,35.37931],[-93.3156175,35.3792996]]]]},"properties":null}`

func TestConvert(t *testing.T) {
	result, err := convert([]byte(operator_acreage_wkt))
	if err != nil {
		t.Fatal(err)
	}
	if got, exp := string(result), operator_acreage_geojson; exp != got {
		t.Fatalf("expected %s, got %s", exp, got)
	}
}