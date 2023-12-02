"use client"

import * as React from "react"
import { MoonIcon, SunIcon } from "@radix-ui/react-icons"
import { useTheme } from "next-themes"

import { Button } from "@/components/ui/button"

export function ModeToggle() {
  const { theme, setTheme, resolvedTheme } = useTheme()

  React.useEffect(() => {
    if (!theme) {
      setTheme('dark')
    }
  }, [theme, setTheme])

  const toggleTheme = () => {

    if (!theme || theme === 'system') {
      setTheme(resolvedTheme === 'dark' ? 'light' : 'dark')
    } else {
      setTheme(theme === 'dark' ? 'light' : 'dark')
    }
  }

  const Icon = !theme || theme === 'dark' || resolvedTheme === 'dark' ? SunIcon : MoonIcon

  return (
    <Button variant="outline" size="icon" onClick={toggleTheme}>
      <Icon className="h-[1.2rem] w-[1.2rem] transition-all" />
      <span className="sr-only">Toggle theme</span>
    </Button>
  )
}
