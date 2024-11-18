'use client'

import { useQuery } from '@tanstack/react-query'

import { getTasks } from '@/http/get-tasks'

import { Table, TableBody, TableCell, TableRow } from '../ui/table'
import { TaskItem } from './task-item'

export function TaskList() {
  const { data } = useQuery({
    queryKey: ['tasks'],
    queryFn: getTasks,
  })

  return (
    <Table>
      <TableBody>
        {data && data.tasks ? (
          data.tasks.map((task) => <TaskItem key={task.id} data={task} />)
        ) : (
          <TableRow>
            <TableCell className="text-center text-base text-muted-foreground">
              No task found
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  )
}
