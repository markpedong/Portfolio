import { TFileItem, deleteFile, getFiles } from '@/api'
import { PRO_TABLE_PROPS } from '@/constants'
import { dateTimeFormatter } from '@/utils'
import { afterModalformFinish } from '@/utils/antd'
import { ActionType, ProColumns, ProTable } from '@ant-design/pro-components'
import { Image, Popconfirm, Typography } from 'antd'
import { useRef } from 'react'

const Files = () => {
  const actionRef = useRef<ActionType>()
  const columns: ProColumns<TFileItem>[] = [
    {
      title: 'Preview',
      align: 'center',
      width: 150,
      render: (_, record) =>
        ['.jpg', '.jpeg', '.png', '.svg', '.webp'].some(ext =>
          record?.file.toLowerCase().endsWith(ext)
        ) ? (
          <Image src={record?.file} alt="image" width={60} height={60} />
        ) : (
          <Typography.Link href={record?.file} target="_blank">
            View
          </Typography.Link>
        )
    },
    {
      title: 'Name',
      align: 'start',
      dataIndex: 'name'
    },
    {
      title: 'File',
      align: 'center',
      dataIndex: 'file'
    },
    {
      title: 'Created At',
      dataIndex: 'created_at',
      align: 'center',
      render: (_, record) => dateTimeFormatter(record.created_at, 'MM-DD-YYYY')
    },
    {
      title: 'Operator',
      align: 'center',
      render: (_, record) => (
        <Popconfirm title="Delete this File?" onConfirm={() => handleDelete(record)}>
          <Typography.Link type="danger">Delete</Typography.Link>
        </Popconfirm>
      )
    }
  ]

  const handleDelete = async (record: TFileItem) => {
    const res = await deleteFile({ url: record?.file })

    return afterModalformFinish(actionRef, res)
  }

  const fetchData = async () => {
    const res = await getFiles()

    return {
      data: res?.data.data ?? [],
      total: res?.data.data?.length ?? 0
    }
  }

  return (
    <ProTable
      {...PRO_TABLE_PROPS}
      rowKey="id"
      columns={columns}
      actionRef={actionRef}
      request={fetchData}
      scroll={{ x: 1100 }}
    />
  )
}

export default Files
