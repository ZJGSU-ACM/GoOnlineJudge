#功能简介
版本14.11.01

###符号说明：

	# 与用户权限相关的方面 

###功能：
	
	总览：

		用户通过导航模块使用不同功能

	功能：

		主页：

			展示新闻、通知，或常用信息、链接等

		问题：

			列表页面：

				列表展示问题的基本信息
				按任题目源、ID、标题等搜索题目
				
				# 权限不同的用户看到的题目列表不同

			问题详情页面：

				显示题目所有的要求和判题信息（描述、通过情况等）
				点击题目源、作者直接转向相关信息
				相关功能，比如提交、通过数据、状态数据
				
		状态：

			展示所有的提交判题状态
			丰富且方便的搜索功能，便于查看自己的状态
			实时刷新，判题状态实时更新
			能通过状态方便的转向其他功能，比如顺便查看题目、用户
			显示代码相似度，可以防作弊

			# 用户权限不同看到的信息不同，如管理员可点击看代码

		竞赛：

			列表页面：

				该模块用于比赛
				展示所有竞赛的基本信息

			竞赛详情页面：

				其中展示竞赛的详细信息
				不同的竞赛由不同的权限（public、private、password）
				竞赛中的所有功能与OJ分开管理，数据显示和提交不共享
				竞赛排名导出，方便教师统计成绩
	
		排名：

			显示用户的成绩以及排名信息
			同样可以方便的搜索

		帮助：

			介绍OJ相关说明和OJ的帮助文档，比如编译环境、判题结果说明、代码示例等


		用户：

			登录页面：


			用户详情页面：

				显示用户的基本信息，包括题目的完成情况以及提交数据排名等
				题目的完成情况包含已解决和未解决的题目，并有提交次数说明
				用户的最近登入纪录

		代码：

			显示用户提交的代码，以及该次提交基本信息
			不同代码不同高亮显示

		管理：

			权限：

				所有管理页面都需要做用户权限的验证
				大致权限从低到高排列如下：
					primary_user			->普通用户
					teacher					->教师
					admin					->超级管理

			功能：

				编辑新闻（添加/删除/修改）
				编辑题目（添加/删除/修改）
				编辑练习（添加/删除/修改）
				编辑权限（添加/删除/修改）
				修改密码
				题目重判
				账号生成
				题目导入
