# dict - 命令行单词查询工具
从爱词霸爬取 **柯林斯词典** 的单词释义和例句, **英文解释优先**

## 为什么写这个工具
个人记单词基本上只会去看 **英文释义** , 因为英文解释能更全面的描述单词的含义. 然而很多单词查询网站都只给出了中文解释, 即便有英文释义(如爱词霸), 也需要向下翻页到中间才能找到，很不方便。

## 使用方法
```
./dict water // 查water单词
```
输出:
```
water - 美 [ˈwɔtɚ, ˈwɑtɚ]
1. [N-UNCOUNT] Water is a clear thin liquid that has no colour or taste when it is pure. It falls from clouds as rain and enters rivers and seas. All animals and people need water in order to live.(水;雨水)
	1. Get me a glass of water.(给我杯水。)
	2. ...the sound of water hammering on the metal roof.(雨水敲打金属屋顶发出的声响)
2. [N-PLURAL] You use waters to refer to a large area of sea, especially the area of sea which is near to a country and which is regarded as belonging to it.(海域;(尤指)领海)
	1. The ship will remain outside Chinese territorial waters.(这艘船将继续呆在中国领海之外。)
	2. ...the open waters of the Arctic Ocean.(北冰洋的公共海域)
3. [N-PLURAL] You sometimes use waters to refer to a situation which is very complex or difficult.(困境)
	1. ...the man brought in to guide him through troubled waters...(被请来帮助他渡过难关的人)
	2. The British Government may be in stormy economic waters.(英国政府可能陷入了动荡的经济困境中。)
4. [VERB] If you water plants, you pour water over them in order to help them to grow.(给(植物)浇水)
	1. He went out to water the plants.(他出去给植物浇水。)
5. [VERB] If your eyes water, tears build up in them because they are hurting or because you are upset.(流泪)
	1. His eyes watered from cigarette smoke.(他的眼睛被香烟熏得直流泪。)
6. [VERB] If you say that your mouth is watering, you mean that you can smell or see some nice food and you might mean that your mouth is producing a liquid.(流口水)
	1. ...cookies to make your mouth water.(令人垂涎的饼干)
7. [PHR-ERG] When a pregnant woman's waters break, the fluid in her womb that surrounds the baby passes out of her body, showing that the baby is ready to be born. A doctor or midwife can break a woman's waters so that the birth can begin.(（孕妇）羊水破裂;（医生或接生员）给（孕妇）破羊水)
	1. My waters broke at six in the morning and within four hours Jamie was born.(我是凌晨6点破水的，不到4个小时杰米就出生了。)
8. [PHRASE] If you say that an event or incident is water under the bridge, you mean that it has happened and cannot now be changed, so there is no point in worrying about it any more.(泼出之水;不可改变的既成事实;无法挽回的过去)
	1. He was relieved his time in jail was over and regarded it as water under the bridge.(他为服刑期满而感到如释重负，而且觉得既已无可更改，过去的就让它过去吧。)
9. [PHRASE] If you are in deep water, you are in a difficult or awkward situation.(陷入困境;处境尴尬)
	1. You certainly seem to be in deep water...(你看起来无疑是陷入了困境。)
	2. I could tell that we were getting off the subject and into deep water.(我觉得我们在跑题并已陷入尴尬境地。)
10. [PHRASE] If an argument or theory does not hold water, it does not seem to be reasonable or be in accordance with the facts.((论点或理论)符合逻辑，站得住脚，经得起考验)
	1. This argument simply cannot hold water in Europe.(这个论点在欧洲根本就站不住脚。)
11. [PHRASE] If you are in hot water, you are in trouble.(陷于困境;遇到麻烦)
	1. The company has already been in hot water over high prices this year.(这家公司今年由于价格过高已经陷入困境。)
12. [PHRASE] If you pour cold water on an idea or suggestion, you show that you have a low opinion of it.(给(想法或建议)泼冷水)
	1. City economists pour cold water on the idea that the economic recovery has begun.(伦敦金融城的经济学家给经济开始复苏的说法泼了冷水。)
13. [PHRASE] If you test the water or test the waters, you try to find out what reaction an action or idea will get before you do it or tell it to people.(试探)
	1. You should be cautious when getting involved and test the water before committing yourself.(参与其中时一定要小心谨慎，表态之前最好试探一下。)
```

## 编译
项目clone下来后, 把项目目录设为`GOPATH`, 然后执行
```
go install dict/client
```

## 正在开发的功能
- dict-server。用于构建单词数据库, 在服务端运行，接收dict-client的查询请求. dict-server从爱词霸爬到结果后写入到本地数据库中, 第二次查该单词时会直接从数据库中读取
- 支持有道等多个网站爬单词
