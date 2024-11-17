import { getTasks } from '@/actions/get-tasks'

import { Table, TableBody } from '../ui/table'
import { TaskItem } from './task-item'

export async function TaskList() {
  const tasks = await getTasks()

  return (
    <Table>
      <TableBody>
        {tasks?.map((task) => <TaskItem key={task.id} data={task} />)}
      </TableBody>
    </Table>
  )
}
