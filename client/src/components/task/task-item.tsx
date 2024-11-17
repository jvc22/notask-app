'use client'

import { Check, Trash } from 'lucide-react'
import { toast } from 'sonner'

import { removeTask } from '@/actions/remove-task'

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
  async function onRemoveTask(text: string) {
    const result = await removeTask({ taskId: id })

    if (result) {
      toast.success(`Task ${text} successfully!`)
    }
  }

  return (
    <TableRow className="text-base">
      <TableCell className="w-1/5">{title}</TableCell>

      <TableCell className="text-muted-foreground">{description}</TableCell>

      <TableCell className="w-1/5">
        <div className="flex justify-end">
          <Button
            size={'sm'}
            variant={'ghost'}
            className="hover:text-green-500"
            onClick={() => onRemoveTask('finished')}
          >
            <Check className="size-3" />
          </Button>

          <Button
            size={'sm'}
            variant={'ghost'}
            className="hover:text-red-500"
            onClick={() => onRemoveTask('deleted')}
          >
            <Trash className="size-3" />
          </Button>
        </div>
      </TableCell>
    </TableRow>
  )
}
