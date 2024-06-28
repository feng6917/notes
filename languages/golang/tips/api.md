##### 编写新接口注意事项

- 需求问题需要严格反复确认
- 需要针对遵守格式的协议字段进行商讨确认
- byte 默认使用base64字符串处理
- int64 默认使用字符串处理

##### 协议定义格式

```
  /版本/公司(项目)/服务/模块/...业务名词/操作
  eg:
  /v1/lol/permission/role/create // 创建角色
  /v1/lol/permission/role/status/update // 修改角色状态

```

- [ ]  定义接口

  - [ ] http

    - [ ] 创建接口 POST

      - [ ] user & user_profile & department_id & commpany_id

      - [ ] 注意点

        ```
        1. 创建及更新数据需要考虑一致性，尽可能使用事务操作
        ```

    - [ ] 更新接口 PUT

      - [ ] update_struct(结构体) & update_filter(更新字段切片)  后续根据需要更新字段中的值映射获取表结构字段

    - [ ] 获取接口 GET

      - [ ] list 获取列表信息

      - [ ] scroll 滚动获取列表信息

        - [ ] 列表信息仅返回基础信息，具体信息通过详情接口获取，但是数据接口应使用信息详情结构避免参数调整

        - [ ] 注意点

          ```
          1. 不到万不得已不要定制接口，同一个接口返回不同的值可使用map形式返回
          ```

      - [ ] describe 获取信息详情

    - [ ] 删除接口 DELETE
