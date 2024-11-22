'use server'

import { HTTPError } from 'ky'
import { cookies } from 'next/headers'
import { z } from 'zod'

import { signInFn } from '@/http/sign-in'
import { signUpFn } from '@/http/sign-up'

const signUpSchema = z
  .object({
    username: z
      .string()
      .min(3, { message: 'Username should have at least 3 characters.' })
      .max(16, { message: 'Username should have a maximum of 40 characters.' }),
    password: z
      .string()
      .min(8, { message: 'Password should have at least 8 characters.' })
      .max(40, { message: 'Password should have a maximum of 40 characters.' }),
    confirm_password: z.string().min(1, { message: 'Confirm your password.' }),
  })
  .refine((data) => data.password === data.confirm_password, {
    path: ['confirm_password'],
    message: 'Passwords must match.',
  })

export async function signUp(data: FormData) {
  const result = signUpSchema.safeParse(Object.fromEntries(data))

  if (!result.success) {
    const errors = result.error.flatten().fieldErrors

    return { success: false, message: null, errors }
  }

  const { username, password } = result.data

  try {
    await signUpFn({
      username,
      password,
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

const signInSchema = z.object({
  username: z.string().min(1, { message: 'Invalid username.' }),
  password: z.string().min(1, { message: 'Invalid password.' }),
})

export async function signIn(data: FormData) {
  const result = signInSchema.safeParse(Object.fromEntries(data))

  if (!result.success) {
    const errors = result.error.flatten().fieldErrors

    return { success: false, message: null, errors }
  }

  const { username, password } = result.data

  try {
    const { token } = await signInFn({
      username,
      password,
    })

    const cookieStore = await cookies()
    cookieStore.set('notask-token', token, {
      path: '/',
      maxAge: 1 * 60 * 60 * 6,
      secure: true,
      sameSite: true,
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
