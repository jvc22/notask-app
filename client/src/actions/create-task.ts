'use server'

import { HTTPError } from 'ky'
import { revalidatePath } from 'next/cache'
import { z } from 'zod'

import { api } from '@/lib/ky'

const createTaskSchema = z.object({
  title: z
    .string()
    .min(1, { message: 'Title should not be empty.' })
    .max(16, { message: 'Title should have a maximum of 16 characters.' }),
  description: z.string().max(40, {message: 'Description should have a maximum of 40 characters.'}).optional(),
})

export async function createTask(data: FormData) {
  const result = createTaskSchema.safeParse(Object.fromEntries(data))

  if (!result.success) {
    const errors = result.error.flatten().fieldErrors

    return { success: false, message: null, errors }
  }

  const { title, description } = result.data

  try {
    await api.post('tasks', {
      json: {
        title,
        description,
      },
    })
  } catch (err) {
    if (err instanceof HTTPError) {
      const { message } = await err.response.json()

      return { success: false, message, errors: null }
    }

    return {
      success: false,
      message: 'Unexpected error. Try again in a few minutes.',
      errors: null,
    }
  }

  revalidatePath('/get-tasks')

  return { success: true, message: null, errors: null }
}
