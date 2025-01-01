import { getLogs, TSessionItem } from '@/api'
import { PRO_TABLE_PROPS } from '@/constants'
import { dateTimeFormatter } from '@/utils'
import { ProColumns, ProTable } from '@ant-design/pro-components'

const Logs = () => {
	const columns: ProColumns<TSessionItem>[] = [
		{
			title: 'Email',
			dataIndex: 'email',
			key: 'email',
			align: 'center'
		},
		{
			title: 'User ID',
			dataIndex: 'user_id',
			key: 'user_id',
			align: 'center'
		},
		{
			title: 'Created At',
			dataIndex: 'created_at',
			key: 'created_at',
			align: 'center',
			render: (_, record) => dateTimeFormatter(record.created_at, 'MM-DD-YYYY HH:MM:ss')
		},
		{
			title: 'Expires At',
			dataIndex: 'expires_at',
			key: 'expires_at',
			align: 'center',
			render: (_, record) => dateTimeFormatter(record.expires_at, 'MM-DD-YYYY HH:MM:ss')
		}
	]

	const fetchLogs = async () => {
		const res = await getLogs()

		return {
			data: res?.data?.data ?? [],
			total: res?.data.data?.length ?? 0
		}
	}

	return <ProTable {...PRO_TABLE_PROPS} rowKey="id" columns={columns} request={fetchLogs} />
}

export default Logs
