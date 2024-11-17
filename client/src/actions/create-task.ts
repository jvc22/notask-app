'use server'

import { HTTPError } from 'ky'
import { z } from 'zod'

import { api } from '@/lib/ky'

const createTaskSchema = z.object({
  title: z
    .string()
    .min(1, { message: 'Title should not be empty.' })
    .max(64, { message: 'Title should have a maximum of 64 characters.' }),
  description: z.string().optional(),
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

  return { success: true, message: null, errors: null }
}
