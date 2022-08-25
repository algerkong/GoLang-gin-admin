const loginRules = {
  username: [
    { required: true, message: '用户名必填', type: 'error', trigger: 'blur' },
    { required: true, message: '用户名必填', type: 'error', trigger: 'change' },
    { whitespace: true, message: '用户名不能为空' },
    { min: 4, message: '输入字数应在3到6之间', type: 'error', trigger: 'blur' },
    { max: 12, message: '输入字数应在3到6之间', type: 'error', trigger: 'blur' },
  ],
  password: [{ required: true, message: '密码必填', type: 'error' }],
}

export { loginRules }
