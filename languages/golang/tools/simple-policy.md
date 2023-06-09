要写一个简单的权限系统，首先要明白权限是什么？
```
    权限，权限是指为了保证职责的有效履行，任职者必须具备的，对某事项进行决策的范围和程度。拆分出来理解就是  某事项（资源）- 程度（增删该查），范围可以作为一种条件来理解。
    譬如：你在网站看到一篇文章，只能看不能编辑，另外该文章设置为订阅才能查看完整内容。
    文章相当于资源；看/编辑相当于决策的程度；内容相当于范围。
``` 
---
 简单了解了权限是什么东西？接下来研究一下怎么写？ 
```   
    譬如：我写文章。
    文章（资源）- 写（程度），就是一条权限。联想一下除了写文章还能干什么？还能读文章，还能写随记。除了我能做些事，张三、李四、王五也能做这些。
    我、张三、李四 作为任职者，可以定义一个名词，譬如用户。
    读、写 作为程度，可以定义一个名词，譬如操作。
    文章、随记 作为资源，可以定义一个名词，譬如资源。
    操作+资源 组成了基本的权限，用户+操作+资源 即组成了一个基本权限系统。
    简单定义一下数据中表结构即；
    users 用户表
    ID int 
    name string
    login_name string
    用户表定义了用户ID、昵称、登陆名；
    actions 操作表
    ID int
    name string
    name_en string
    操作表定义了操作ID、名称、英文名；
    resource 资源表
    ID int
    name string
    name_en string
    资源表定义了操作ID、名称、英文名；
    asso 用户操作资源关联表
    user_id int
    action_id int
    resource_id int
    用户操作资源定义了操作ID、资源ID，将用户与操作和资源关联起来；一套简单的权限系统就出来了。
``` 
---
再复杂一点，如果是用户组呢？
```
    用户+操作+资源 只是基本的权限系统，任职者可能不仅仅是用户，如果是用户所在的部门怎么办？这里就引入了角色的概念，用户属于哪个角色，角色可以为部门，可以为管理员，可以为职位，角色可以理解为多个用户的集合。
    user+role+action+resource 组成了一套新的权限系统。
    读文章可能不仅仅是张三需要，李四也需要，王五也需要怎么办？这里把操作和资源组合起来，定义为权限，张三有读文章的权限，李四、王五也有。
    user+role+permission(action+resource) 组成了一套新的权限系统。     
``` 

再复杂一点，权限有先后顺序怎么办呢？
```
    如果同一个角色有查看、编辑、删除权限，我想把查看权限放在最后，删除权限放在最前面展示怎么办？
    这里引入了排序的定义，可以在关联关系中增加排序字段。    
    user+role+permission(action+resource):sort
    如果拥有编辑权限，但是没有查看权限却想要默认拥有怎么办？
    user+role+permission(action+resource):sort:isDefault
    如果还有其他，更多更多筛选条件呢？可以定义一个condition字段，类型设置为json，用来做筛选或特殊情况处理。
    user+role+permission(action+resource):condition
    
``` 
---
总结一下行不行的通。
```
    user+role+permission(action+resource):condition
    定义用户表
    定义角色表
    定义用户角色关联表
    定义操作表
    定义资源表
    定义操作资源组合表-权限表
    定义用户/角色与权限表关联表
    这样就可以实现了一套完整的权限系统。
    另外一般情况下，都是开发或管理员设定好相应权限，由客户创建角色及绑定权限，权限控制在一定范围内，这样来说，权限是在服务启动那一刻已经定死的，可以定义一个初始表定义好相关字段，
用户绑定权限从其中选择。
    另外一个服务下可能区分应用的概念，一个应用可以自己应用及其他应用下的权限进行绑定，可以定义一个应用权限使用表，定义父级及子级应用ID来处理应用权限使用关系。    
```


<img width="639" alt="image" src="https://user-images.githubusercontent.com/82997695/210159573-21bec2da-adb1-45b3-8c94-675b565a3e9f.png">