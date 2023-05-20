# 如何安装和使用自主人工智能工具Auto-GPT



## 前言

🔮 在我的 Slack 工作区中，集成了多个 AI，分别有 ChatGPT 4、ChatGPT 3.5、Claude ……

我们可以通过 Slack 免费并且无限制的和 AI 交互，欢迎大家加入到 Slack，这里是 链接：

[https://join.slack.com/t/kubecub/shared_invite/zt-1se0k2bae-lkYzz0_T~BYh3rjkvlcUqQ](https://join.slack.com/t/kubecub/shared_invite/zt-1se0k2bae-lkYzz0_T~BYh3rjkvlcUqQ)

![image-20230514215132365](http://sm.nsddd.top/sm202305142151717.png)



## 介绍

很早之前就了解到了 Auto-GPT，作为 GitHub 上近期增长速度最快的项目（没有之一），Auto-GPT 在开源社区可谓是人尽皆知，甚至 star 已经很快就超过 Kubernetes，目前有 `125k star`。

得益于 Auto-GPT 的出色技术，可以高精度和高效率地自动执行许多任务。 它利用了 GPT-4 强大的自然语言处理功能。

我们甚至可以通过它来实现更多的自动化的工作，比如说前一节 [在 Sealos 上开发一款 AI 自动云原生化项目自动上线工具](./49.md)



## 什么是 AutoGPT

它的 GitHub 地址：

+ [GitHub](https://github.com/Significant-Gravitas/Auto-GPT)

从本质上讲，Auto-GPT 利用 OpenAI 最新人工智能模型的多功能性与软件和服务进行在线互动，使其能够 “自主 “执行X和Y等任务。但正如我们在大型语言模型方面的学习，这种能力似乎像海洋一样宽广，但却像水坑一样深。

AutoGPT 是一个由人工智能驱动的应用程序，利用 GPT-4 等 LLM 的强大功能自主创建和处理各种工作。通过使用 Auto GPT，组织和个人可以简化报告创作、内容创建和数据分析等流程，以节省时间并减少错误。

AutoGPT 改变了任务自动化的游戏规则，使组织和个人能够专注于其他关键任务，同时将重复和琐碎的工作留给程序。

随着 LLM 的不断发展，我们可以期待看到像 Auto GPT 这样功能越来越强大的软件能够执行越来越复杂的任务。

就人工智能驱动的技术将如何改变我们未来的操作方式和与人工智能系统的互动方式而言，AutoGPT 提供了一个新方向。



## 它是如何工作的

Auto GPT 使用 LLM 的最新发展，特别是 GPT-4，自动生成具有凝聚力和相关性的内容。该程序从大量的数据中学习，这使它能够识别单词和句子之间的模式和联系。

使用这些信息，Auto GPT 会生成文本以响应提示或输入。这种输入可能以指令、任务或一套指导方针的形式出现。

Auto GPT 在收到输入后，使用其尖端的算法和自然语言处理技能，创建上下文适当且逻辑一致的内容。对于希望自动化流程和节省时间的组织和人员来说，Auto GPT 是一种重要的资源，因为它生成的文本几乎与人类书面语言无法区分。

Auto GPT 的优势在于它能够从大量数据中学习，并生成相关和合乎逻辑的文本，使其成为作业自动化领域的关键工具。

与免费版的ChatGPT不同，**Auto-GPT可以连接到互联网**，找到任何主题的最新信息。因此，你可以用它来访问任何网页并捕获信息。



### 可以做哪些事情

Auto GPT是一个灵活的程序，可用于各种活动，包括创建报告和数据分析。在这一部分中，我们将了解 Auto GPT可以执行的一些功能，以及它是如何自动执行这些功能的。

**内容创建**

网站、博客和社交媒体帖子的内容可以使用 Auto GPT 创建。如果你给它一个主题或一套指导方针，Auto GPT 可以产生高质量、相关性和趣味性的材料。

**翻译**

你可以使用 Auto GPT 执行翻译任务。通过使用 Auto GPT 以一种语言作为输入文本，可以将文本翻译成另一种语言。在不同国家/地区开展业务并需要快速文档或通信翻译的企业可以提供非常大帮助。

**客户服务**

客户支持职责，如响应频繁的查询和解决问题，可以通过 Auto GPT 实现自动化。Auto GPT 可以使用自然语言处理来理解客户查询并提供相关的解决方案。

**数据分析**

可以使用 Auto GPT 执行数据分析活动。数据输入允许 Auto GPT 分析信息并产生可用于决策的见解。

**撰写报告**

企业和研究人员可以使用 Auto GPT 根据数据输入生成报告。通过输入数据，Auto GPT可以分析信息并产生准确和有指导意义的结果。

**编码**

Auto GPT 可用于编码作业生成完整的程序或代码片段。Auto GPT 可以通过考虑编程参数或需求来生成有效且高效的代码。需要精确快速地编写代码的开发人员会发现这种功能非常有帮助。

*发挥你的想象力，只有你想不到的，没有它做不到的~*



## 搭建和设置环境

如果你可以**[使用GPT-4](https://www.wbolt.com/how-to-use-gpt-4-free.html) API**，Auto-GPT的效果最好，因为它更善于思考并得出结论。它也不太容易产生幻觉。如果你还没有访问权限，你可以[使用这里的链接](https://www.wbolt.com/go?_=18a71268abaHR0cHM6Ly9vcGVuYWkuY29tL3dhaXRsaXN0L2dwdC00LWFwaQ%3D%3D)加入GPT-4 API访问的等待名单。然而，你也可以使用普通的OpenAI API与GPT-3.5模型。

**准备工作：**

1. Git

2. Python 3.8 or later

3. OpenAI API key

4. PineCone API key



用 git clone:

```bash
❯ git clone https://github.com/Significant-Gravitas/Auto-GPT
❯ cd Auto-GPT
```

安装：

> requirements.txt 文件是一个文本文件，通常用于 Python 项目中，其中包含了该项目所需的所有依赖包及其版本信息。
>
> 我们之前在学习 buildpacks 项目的时候，它如何解决 python 的环境判断问题，就是通过 requirements.txt 文件。

```bash
❯ pip install -r requirements.txt
```

之后，将`.env.template` 重命名为 `.env`，并使用 OpenAI 和 PineCone API 密钥填充字段。

```bash
❯ mv .env.template .env
```

之后，转到 VIM ，在 `OPENAI_API_KEY` 部分粘贴API。你可以参考下面的图片来了解一下。

```bash
❯ cat .env | grep -i OPENAI_API_KEY
## OPENAI_API_KEY - OpenAI API Key (Example: my-openai-api-key)
OPENAI_API_KEY=your-openai-api-key
```



接下来，打开[pinecone.io](https://www.wbolt.com/go?_=baa37b74edaHR0cHM6Ly93d3cucGluZWNvbmUuaW8v)并创建一个免费账户。它将允许LLM从内存中检索相关信息，用于AI应用。

在这里，点击左侧边栏的 `API Keys`，并点击右侧窗格的 `Create API Key`

给一个名字，如 `autogpt`，然后点击 `Create Key`

![image-20230503171039991](http://sm.nsddd.top/sm202305031712718.png)

复制 Key Value，用 `vim` 打开，将其粘贴在 `PINECONE_API_KEY` 旁边。

同样地，复制 `Environment` 下的数值。

现在，把它粘贴到`PINECONE_ENV`旁边。



## 运行

打开一个终端来执行 `main.py` Python 脚本。

```bash
❯ python3 -m autogpt
```

**为 AI 命名：**

在第一次运行时，Auto-GPT会要求你**为AI命名**。如果你不想为一个特定的用例创建一个人工智能，你可以把这个字段留空并点击回车。它默认加载的是Entrepreneur-GPT名称。

**定义人工智能的作用~**

之后，**为自主AI逐一设定目标**。这是你告诉人工智能你想实现的目标的地方。你可以要求它将信息保存在一个文本或PDF文件中。你也可以要求它在检索完所有信息后关闭。

**定义 AI 角色：**

根据你希望 AI 发挥的功能，为其命名和角色，例如“研究人员”、“内容生成器”或“个人编码器”。为了获得更成功的结果，明确你希望人工智能实现的目标。

**设定目标：**

详细概述人工智能的目标，例如获取信息、将数据存储在文件中、执行代码或修改文本。包括要使用的输出文件的信息，以及完成作业所需的任何其他操作。

目标如下：

1. 开发产品

2. 优化产品

3. 扩大产品规模

4. 创造5000万美元以上的收入

5. 保持这种一致性

Auto-GPT将开始思考。在行动过程中，它将要求你授权行动。**按 “y”**，然后按**回车键**确认。它可能会连接到网站并收集信息。

你可以读懂**人工智能在思考、推理和计划什么**。它还提供批评（一种负面的提示），以便它拿出正确的信息。最后，它执行行动。



## AutoGPT 插件

你可以根据自己的独特需求调整 AutoGPT

它们不需要对主要应用程序的核心代码进行重大更改，因为它们是为了扩展或改进其功能而设计的。

插件列表如下：

1、Twitter plugin

2、Email plugin

3、Telegram plugin

4、Google Analytics plugin

5、Youtube plugin, and many more.



## Auto GPT 和 LLM 的未来

尽管像 GPT-4 这样的 LLM 为工作自动化的革新提供了很大的希望，但也可能存在需要考虑的危险和缺点。用于训练模型的数据中可能存在的偏见和成见是令人担忧的关键原因之一。有偏见的LLM可能会出现不公平和歧视性的结果。

正如 Auto-GPT 和 ChatGPT 所展示的那样，可以教导 LLM 从巨大的数据量中学习，并独立开展广泛的活动，从内容制作到编码。自动化操作有能力完全改变行业和我们的运作方式。

但对于 LLM 来说，Auto-GPT 只是一个开始。随着技术的进一步发展，LLM 的权力将会增加。未来的 LLM 将更善于完成复杂的任务，理解背景和复杂性。

LLM 任务自动化也可能开辟新的市场和就业机会。如果企业和人们能够将许多平凡的琐事自动化，他们将能够专注于更困难和富有想象力的项目。



## Auto-GPT替代品： 用AgentGPT实现任务自动化

如果你不想在本地设置Auto-GPT，并希望有一个易于使用的解决方案来自动化和部署任务，你可以使用AgentGPT。它建立在Auto-GPT上，但你可以在浏览器中直接访问它。不需要摆弄终端和命令。以下是它的工作原理。



### 打开使用

+ [https://agentgpt.reworkd.ai/zh](https://agentgpt.reworkd.ai/zh)

在这里，添加你的OpenAI API密钥。你可以从[这里](https://www.wbolt.com/go?_=e7e2fef44aaHR0cHM6Ly9wbGF0Zm9ybS5vcGVuYWkuY29tL2FjY291bnQvYXBpLWtleXM%3D)**获得API密钥**。如果你不能访问[GPT-4](https://www.wbolt.com/how-to-use-gpt-4-free.html) API，选择 `gpt-3.5-turbo` 作为模型，然后点击 `Save`

![image-20230503174431508](http://sm.nsddd.top/sm202305031744911.png)

接下来，给你的人工智能代理起个名字，并设定你希望实现的目标。现在，点击Auto-GPT AI的 `Deploy Agent`，开始考虑你的投入。

我发现我没钱了~

![image-20230503175554002](http://sm.nsddd.top/sm202305031755504.png)

一旦任务完成，你可以点击 “**Save**“或 “Copy” 来获得最终结果。



## END 链接

### 参考文章

+ [如何安装和使用自主人工智能工具Auto-GPT](https://www.wbolt.com/auto-gpt.html)
+ [使用 Auto-GPT 的八种方法](https://juejin.cn/post/7226665757882449978)
+ [知乎：安装和使用 Auto-GPT](https://zhuanlan.zhihu.com/p/621854926)
+ [Auto-GPT-ZH 自主 GPT4 实验](https://github.com/kaqijiang/Auto-GPT-ZH)
+ [高级 LLMOps 架构](https://matt-rickard.com/a-high-level-llmops-architecture)

