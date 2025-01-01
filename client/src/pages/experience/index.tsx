import {
  TExperienceItem,
  addExperiences,
  deleteExperiences,
  getExperiences,
  toggleExperienceStatus,
  updateExperiences
} from '@/api'
import { GLOBAL_STATUS } from '@/api/constants'
import { MODAL_FORM_PROPS, PRO_TABLE_PROPS } from '@/constants'
import { capFrstLtr, randomColorGenerator } from '@/constants/helper'
import { dateTimeFormatter } from '@/utils'
import { afterModalformFinish } from '@/utils/antd'
import {
  ActionType,
  ModalForm,
  ProColumns,
  ProForm,
  ProFormDatePicker,
  ProFormList,
  ProFormSlider,
  ProFormText,
  ProTable
} from '@ant-design/pro-components'
import { Button, Popconfirm, Space, Switch, Tag, Typography } from 'antd'
import { useRef } from 'react'

const Experience = () => {
  const actionRef = useRef<ActionType>()
  const columns: ProColumns<TExperienceItem>[] = [
    {
      title: 'Title',
      align: 'center',
      dataIndex: 'title'
    },
    {
      title: 'Company',
      align: 'start',
      dataIndex: 'company'
    },
    {
      title: 'Location',
      align: 'center',
      dataIndex: 'location'
    },
    {
      title: (
        <div className="flex flex-col gap-0">
          <div>Started</div>
          <div>Ended</div>
        </div>
      ),
      align: 'center',
      width: 100,
      search: false,
      render: (_, record) => (
        <div className="flex flex-col">
          <div>{record.started}</div>
          <div>{record.ended}</div>
        </div>
      )
    },
    {
      title: 'Descriptions',
      align: 'start',
      search: false,
      ellipsis: true,
      width: 200,
      render: (_, record) =>
        record?.descriptions?.map(q => (
          <span className="overflow-hidden text-ellipsis" key={q}>
            - {q} <br />
          </span>
        ))
    },
    {
      title: 'Tech Stack',
      align: 'left',
      search: false,
      render: (_, record) => (
        <Space direction="vertical" align="start">
          {record?.skills?.map(q => (
            <Tag key={q?.id} color={`#${randomColorGenerator()}`}>
              {capFrstLtr(q?.name)} - {q?.percentage}%
            </Tag>
          ))}
        </Space>
      )
    },
    {
      title: (
        <div className="flex flex-col gap-0">
          <div>Created</div>
          <div>Updated</div>
        </div>
      ),
      search: false,
      align: 'center',
      width: 160,
      render: (_, record) => (
        <div className="flex flex-col">
          <div>{dateTimeFormatter(record.created_at, 'MM-DD-YYYY HH:MM:ss')}</div>
          <div>{dateTimeFormatter(record.updated_at, 'MM-DD-YYYY HH:MM:ss')}</div>
        </div>
      )
    },
    {
      title: 'Operator',
      align: 'center',
      width: 180,
      search: false,
      render: (_, record) => (
        <Space>
          {renderSwitch(record)}
          {renderAddEditExperience('EDIT', record)}
          {renderDeleteExperience(record)}
        </Space>
      )
    }
  ]

  const renderSwitch = (record: TExperienceItem) => {
    return (
      <Switch
        unCheckedChildren="OFF"
        checkedChildren="ON"
        checked={record?.status === GLOBAL_STATUS.ON}
        onChange={async () => {
          const res = await toggleExperienceStatus({ id: record?.id })

          return afterModalformFinish(actionRef, res)
        }}
      />
    )
  }

  const renderAddEditExperience = (type: 'ADD' | 'EDIT', record?: TExperienceItem) => {
    const isEdit = type === 'EDIT'
    return (
      <ModalForm
        {...MODAL_FORM_PROPS}
        initialValues={
          isEdit
            ? { ...record, descriptions: record?.descriptions?.map(q => ({ description: q })) }
            : {}
        }
        title={isEdit ? 'Edit Experience' : 'Add Experience'}
        trigger={
          isEdit ? <Typography.Link>Edit</Typography.Link> : <Button type="primary">ADD</Button>
        }
        grid
        layout="inline"
        rowProps={{
          gutter: [0, 25]
        }}
        width={1200}
        onFinish={async params => {
          let res

          if (isEdit) {
            res = await updateExperiences({
              ...params,
              descriptions: params?.descriptions?.map(q => q?.description),
              id: record?.id
            })
          } else {
            res = await addExperiences({
              ...params,
              descriptions: params?.descriptions?.map(q => q?.description)
            })
          }

          return afterModalformFinish(actionRef, res)
        }}
      >
        <ProForm.Group>
          <ProForm.Group colProps={{ span: 9 }} rowProps={{ gutter: [0, 20] }}>
            <ProFormText label="Title" name="title" rules={[{ required: true }]} />
            <ProFormText label="Company" name="company" rules={[{ required: true }]} />
            <ProFormText label="Location" name="location" rules={[{ required: true }]} />
            <ProFormDatePicker label="Started" name="started" rules={[{ required: true }]} />
            <ProFormDatePicker label="Ended" name="ended" rules={[{ required: true }]} />
          </ProForm.Group>
          <ProForm.Group colProps={{ span: 15 }}>
            <ProFormList
              name="descriptions"
              label="Descriptions"
              creatorButtonProps={{
                position: 'bottom',
                creatorButtonText: 'Add Description'
              }}
              copyIconProps={false}
              alwaysShowItemLabel
              style={{ marginBottom: '1rem' }}
            >
              <div style={{ marginBottom: '1rem' }}>
                <ProFormText
                  name="description"
                  label="Description"
                  labelCol={{ flex: '100px' }}
                  rules={[{ required: true }]}
                />
              </div>
            </ProFormList>
            <ProFormList
              name="skills"
              label="Tech Stack"
              creatorButtonProps={{
                position: 'bottom',
                creatorButtonText: 'Add Skill'
              }}
              copyIconProps={false}
              alwaysShowItemLabel
            >
              <ProForm.Group style={{ marginBottom: '1rem' }}>
                <ProFormText
                  name="name"
                  label="Technology"
                  colProps={{ span: 12 }}
                  labelCol={{ flex: '100px' }}
                  rules={[{ required: true }]}
                />
                <ProFormSlider
                  name="percentage"
                  label="Percentage"
                  colProps={{ span: 12 }}
                  labelCol={{ flex: '100px' }}
                  rules={[{ required: true }]}
                />
              </ProForm.Group>
            </ProFormList>
          </ProForm.Group>
        </ProForm.Group>
      </ModalForm>
    )
  }
  const renderDeleteExperience = (record: TExperienceItem) => {
    return (
      <Popconfirm
        title="Delete this Experience?"
        onConfirm={async () => {
          const res = await deleteExperiences({ id: record?.id })

          return afterModalformFinish(actionRef, res)
        }}
      >
        <Typography.Link type="danger">Delete</Typography.Link>
      </Popconfirm>
    )
  }

  const fetchData = async () => {
    const res = await getExperiences()

    return {
      data: res?.data.data
    }
  }

  return (
    <div>
      <ProTable
        {...PRO_TABLE_PROPS}
        rowKey="id"
        columns={columns}
        actionRef={actionRef}
        request={fetchData}
        toolBarRender={() => [renderAddEditExperience('ADD')]}
        scroll={{ x: 1000 }}
      />
    </div>
  )
}

export default Experience
