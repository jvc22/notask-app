'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { useMutation } from '@tanstack/react-query'
import { AlertTriangle } from 'lucide-react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { createTask } from '@/http/create-task'
import { getQueryClient } from '@/lib/react-query'
import { cn } from '@/lib/utils'

const createTaskSchema = z.object({
  title: z
    .string()
    .min(1, { message: 'Title should not be empty.' })
    .max(20, { message: 'Title should have a maximum of 20 characters.' }),
  description: z
    .string()
    .max(40, { message: 'Description should have a maximum of 40 characters.' })
    .optional(),
})

type CreateTaskData = z.infer<typeof createTaskSchema>

export function NewTaskForm() {
  const queryClient = getQueryClient()

  const {
    register,
    reset,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<CreateTaskData>({
    resolver: zodResolver(createTaskSchema),
  })

  const { mutateAsync: createTaskFn, isError } = useMutation({
    mutationFn: createTask,
    onSuccess: async () => {
      reset()

      queryClient.invalidateQueries({
        queryKey: ['tasks'],
      })

      toast.success('Task created successfully!')
    },
  })

  return (
    <div className="space-y-3">
      {isError && (
        <Alert className="border-red-500 text-red-500">
          <AlertTriangle className="size-4 stroke-red-500" />
          <AlertTitle>Submission failed.</AlertTitle>
          <AlertDescription>
            Unexpected error. Try again in a few minutes.
          </AlertDescription>
        </Alert>
      )}

      <form
        onSubmit={handleSubmit((data) => createTaskFn(data))}
        className="mx-md:grid-rows-3 grid gap-3 md:grid-cols-7"
      >
        <Input
          {...register('title')}
          autoComplete="off"
          placeholder="Task title"
          className={cn(
            'dark:bg-zinc-900 dark:shadow-shape md:col-span-2',
            errors?.title ? 'border-red-500' : 'dark:border-transparent',
          )}
        />

        <Input
          {...register('description')}
          autoComplete="off"
          placeholder="Task description (optional)"
          className="dark:border-transparent dark:bg-zinc-900 dark:shadow-shape md:col-span-4"
        />

        <Button
          type="submit"
          disabled={isSubmitting}
          className="dark:bg-zinc-800 dark:text-foreground dark:shadow-shape dark:hover:bg-zinc-800/90"
        >
          Create
        </Button>
      </form>

      {errors?.title ? (
        <div>
          <span className="text-sm font-medium text-red-500">
            {errors.title.message}
          </span>
        </div>
      ) : (
        errors?.description && (
          <div>
            <span className="text-sm font-medium text-red-500">
              {errors.description.message}
            </span>
          </div>
        )
      )}
    </div>
  )
}
