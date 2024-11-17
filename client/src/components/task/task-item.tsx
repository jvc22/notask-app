import { Check, Trash } from 'lucide-react'

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
  return (
    <TableRow className="text-base">
      <TableCell className="w-1/5">{title}</TableCell>

      <TableCell className="text-muted-foreground truncate">
        <span className="truncate">{description}</span>
      </TableCell>

      <TableCell className="w-1/5">
        <div className="flex justify-end">
          <Button
            size={'sm'}
            variant={'ghost'}
            className="hover:text-green-500"
          >
            <Check className="size-3" />
          </Button>

          <Button size={'sm'} variant={'ghost'} className="hover:text-red-500">
            <Trash className="size-3" />
          </Button>
        </div>
      </TableCell>
    </TableRow>
  )
}
