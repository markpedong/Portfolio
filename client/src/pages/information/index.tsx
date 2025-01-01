import { getInfo, updateDetails } from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { FORM_PROPS, INPUT_EMAIL } from '@/constants'
import { INPUT_TRIM } from '@/utils'
import { afterModalformFinish } from '@/utils/antd'
import {
  ProForm,
  ProFormInstance,
  ProFormSwitch,
  ProFormText,
  ProFormTextArea,
  ProFormUploadButton
} from '@ant-design/pro-components'
import { Typography } from 'antd'
import { omit } from 'lodash'
import { useRef } from 'react'

const Information = () => {
  const formRef = useRef<ProFormInstance>()
  const getValue = (key: string) => formRef?.current?.getFieldValue(key)

  const fetchData = async () => {
    const res = await getInfo()

    return {
      ...res?.data.data,
      resume_pdf: [{ url: res?.data.data.resume_pdf, name: res?.data.data.resume_pdf }],
      resume_docx: [{ url: res?.data.data.resume_docx, name: res?.data.data.resume_docx }]
    }
  }

  return (
    <div>
      <Typography.Title level={5}>Update the main details of the website: </Typography.Title>
      <ProForm
        {...FORM_PROPS}
        labelCol={{ flex: '120px' }}
        formRef={formRef}
        grid
        request={fetchData}
        onFinish={async params => {
          const res = await updateDetails(
            omit(
              {
                ...params,
                id: getValue('id'),
                username: getValue('username'),
                resume_pdf: getValue('resume_pdf')?.[0]?.url,
                resume_docx: getValue('resume_docx')?.[0]?.url,
                isdownloadable: params.isdownloadable ? 1 : 0
              },
              ['password']
            )
          )

          fetchData()
          return afterModalformFinish(undefined, res, formRef)
        }}
      >
        <ProFormText
          {...INPUT_TRIM}
          rules={[{ required: true }]}
          label="First Name"
          name="first_name"
          colProps={{ span: 12 }}
        />
        <ProFormText
          {...INPUT_TRIM}
          rules={[{ required: true }]}
          label="Last Name"
          name="last_name"
          colProps={{ span: 12 }}
        />
        <ProFormText
          {...INPUT_TRIM}
          rules={[{ required: true }]}
          label="Phone"
          name="phone"
          colProps={{ span: 12 }}
        />
        <ProFormText
          {...INPUT_TRIM}
          rules={[INPUT_EMAIL, { required: true }]}
          label="Email"
          name="email"
          colProps={{ span: 12 }}
        />
        <ProFormTextArea rules={[{ required: true }]} label="Address" name="address" />
        <ProFormTextArea rules={[{ required: true }]} label="Description" name="description" />
        <ProFormSwitch
          label="Resumé"
          name="isdownloadable"
          checkedChildren={
            <span className="font-bold" style={{ letterSpacing: '0.2rem', fontWeight: 800 }}>
              ENABLED
            </span>
          }
          unCheckedChildren={
            <span className="font-bold" style={{ letterSpacing: '0.2rem', fontWeight: 800 }}>
              DISABLED
            </span>
          }
          fieldProps={{ value: getValue('isdownloadable') === GLOBAL_STATUS.ON ? true : false }}
        />
        <ProFormUploadButton
          label="Resumé PDF"
          name="resume_pdf"
          colProps={{ span: 12 }}
          rules={[{ required: true }]}
        />
        <ProFormUploadButton
          label="Resumé DOCX"
          name="resume_docx"
          colProps={{ span: 12 }}
          rules={[{ required: true }]}
        />
      </ProForm>
    </div>
  )
}

export default Information
