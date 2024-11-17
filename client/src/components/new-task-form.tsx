'use client'

import { AlertTriangle } from 'lucide-react'
import { toast } from 'sonner'

import { createTask } from '@/actions/create-task'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { useFormState } from '@/hooks/use-form-state'
import { cn } from '@/lib/utils'

import { Alert, AlertDescription, AlertTitle } from './ui/alert'

export function NewTaskForm() {
  const [{ success, message, errors }, handleSubmit, isPending] = useFormState(
    createTask,
    () => {
      toast.success('Task created successfully!')
    },
    true,
  )

  return (
    <div className="space-y-3">
      {!success && message && (
        <Alert className="border-red-400 text-red-400">
          <AlertTriangle className="size-4 stroke-red-400" />
          <AlertTitle>Registration failed.</AlertTitle>
          <AlertDescription>{message}</AlertDescription>
        </Alert>
      )}

      <form
        onSubmit={handleSubmit}
        className="mx-md:grid-rows-3 grid gap-3 md:grid-cols-7"
      >
        <Input
          name="title"
          placeholder="Task title"
          className={cn(
            'md:col-span-2 dark:bg-zinc-900 dark:shadow-shape',
            errors?.title ? 'border-red-400' : 'dark:border-transparent',
          )}
        />

        <Input
          name="description"
          placeholder="Task description (optional)"
          className="md:col-span-4 dark:border-transparent dark:bg-zinc-900 dark:shadow-shape"
        />

        <Button
          type="submit"
          className="dark:text-foreground dark:bg-zinc-800 dark:shadow-shape dark:hover:bg-zinc-800/90"
          disabled={isPending}
        >
          Create
        </Button>
      </form>
    </div>
  )
}
