import { useMutation } from '@tanstack/react-query'
import { X } from 'lucide-react'
import { toast } from 'sonner'

import { GetTasksResponse } from '@/http/get-tasks'
import { removeTask } from '@/http/remove-task'
import { getQueryClient } from '@/lib/react-query'

import { Button } from '../ui/button'
import { TableCell, TableRow } from '../ui/table'

interface TaskItemProps {
  data: {
    id: number
    title: string
    description: string
  }
}

export function TaskItem({ data: { id, title, description } }: TaskItemProps) {
  const queryClient = getQueryClient()

  function removeTaskOnCache() {
    const tasksCache = queryClient.getQueriesData<GetTasksResponse>({
      queryKey: ['tasks'],
    })

    tasksCache.forEach(([cacheKey, cached]) => {
      if (!cached) {
        return
      }

      queryClient.setQueryData<GetTasksResponse>(cacheKey, {
        ...cached,
        tasks: cached.tasks.filter((task) => task.id !== id),
      })
    })
  }

  const { mutateAsync: removeTaskFn } = useMutation({
    mutationFn: removeTask,
    onSuccess: async () => {
      removeTaskOnCache()

      toast.success('Task closed successfully!')
    },
  })

  return (
    <TableRow className="text-base">
      <TableCell className="w-[200px] truncate">
        <span title={title}>{title}</span>
      </TableCell>

      <TableCell className="text-muted-foreground">{description}</TableCell>

      <TableCell className="w-[64px]">
        <div className="flex justify-end">
          <Button
            size={'sm'}
            variant={'ghost'}
            onClick={() => removeTaskFn({ taskId: id })}
          >
            <X className="size-3" />
          </Button>
        </div>
      </TableCell>
    </TableRow>
  )
}
