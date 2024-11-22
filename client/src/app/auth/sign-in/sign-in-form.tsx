'use client'

import { AlertTriangle } from 'lucide-react'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'

import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { useFormState } from '@/hooks/use-form-state'
import { cn } from '@/lib/utils'

import { signIn } from '../actions'

export function SignInForm() {
  const router = useRouter()

  const [{ success, message, errors }, handleSubmit, isPending] = useFormState(
    signIn,
    () => {
      router.push('/')

      toast.success('Welcome!')
    },
    false,
  )

  return (
    <form onSubmit={handleSubmit} className="flex flex-col items-center gap-6">
      <div className="text-center">
        <h1 className="text-3xl font-bold tracking-tight">
          notask<span className="text-muted-foreground">.app</span>
        </h1>

        <span className="text-muted-foreground">
          Enjoy your time and improve your productivity.
        </span>
      </div>

      {!success && message && (
        <Alert className="border-red-500 text-red-500">
          <AlertTriangle className="size-4 stroke-red-500" />
          <AlertTitle>Sing in failed.</AlertTitle>
          <AlertDescription>{message}</AlertDescription>
        </Alert>
      )}

      <div className="w-full space-y-3">
        <div className="space-y-1.5">
          <Label htmlFor="username">Username</Label>
          <Input
            id="username"
            name="username"
            autoComplete="off"
            placeholder="Username"
            className={cn(errors?.username && 'border-red-500')}
          />

          {errors?.username && (
            <div>
              <span className="text-sm font-medium text-red-500">
                {errors.username[0]}
              </span>
            </div>
          )}
        </div>

        <div className="space-y-1.5">
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            name="password"
            type="password"
            autoComplete="off"
            placeholder="••••••••"
            className={cn(errors?.password && 'border-red-500')}
          />

          {errors?.password && (
            <div>
              <span className="text-sm font-medium text-red-500">
                {errors.password[0]}
              </span>
            </div>
          )}
        </div>
      </div>

      <div className="flex items-center gap-3">
        <Button
          type="button"
          variant={'link'}
          className="px-0"
          disabled={isPending}
          onClick={() => router.push('/auth/sign-up')}
        >
          Create new account
        </Button>

        <span className="text-muted-foreground">or</span>

        <Button
          type="submit"
          disabled={isPending}
          className="dark:bg-zinc-800 dark:text-foreground dark:shadow-shape dark:hover:bg-zinc-800/90"
        >
          Sign in
        </Button>
      </div>
    </form>
  )
}
