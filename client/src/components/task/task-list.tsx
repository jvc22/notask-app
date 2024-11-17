import { getTasks } from '@/actions/get-tasks'

import { Table, TableBody, TableCell, TableRow } from '../ui/table'
import { TaskItem } from './task-item'

export async function TaskList() {
  const tasks = await getTasks()

  return (
    <Table>
      <TableBody>
        {tasks ? (
          tasks.map((task) => <TaskItem key={task.id} data={task} />)
        ) : (
          <TableRow>
            <TableCell className="text-muted-foreground text-center text-base">
              No tasks found
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  )
}
