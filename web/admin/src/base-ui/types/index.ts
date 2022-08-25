type IFormType = 'input' | 'select' | 'checkbox' | 'datepicker'

export interface IFormItem {
  field: string
  type: IFormType
  label: string
  placeholder?: string
  options?: any[]
  outherOptions?: any
  style?: any
  isHidden?: boolean
}

export interface IForm {
  formItems: IFormItem[]
  labelWidth?: string
  itemStyle?: any
  colLayout?: any
  rules?: any
}
