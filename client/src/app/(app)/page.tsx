import { Header } from '@/components/header'
import { TaskList } from '@/components/task/task-list'

import { NewTaskForm } from './new-task-form'

export default function Home() {
  return (
    <div className="mx-auto w-full max-w-2xl space-y-6">
      <Header />

      <NewTaskForm />

      <TaskList />
    </div>
  )
}
