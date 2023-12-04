'use client'

import * as React from 'react'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'
import { Table, TableHeader, TableRow, TableHead, TableBody, TableCell } from '@/components/ui/table'
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from '@/components/ui/collapsible'
import { ChevronsUpDown } from 'lucide-react'
import { KmeansData } from '@/app/page'
import { Button } from './ui/button'

export default function KmeansClusters({ kmeans }: { kmeans: KmeansData }) {
  const [openClusters, setOpenClusters] = React.useState<Record<string, boolean>>(() => {
    const initialOpenClusters: Record<string, boolean> = {}
    Object.keys(kmeans).forEach((clusterName) => {
      initialOpenClusters[clusterName] = false
    })
    return initialOpenClusters
  })

  const handleToggle = (clusterName: string) => {
    setOpenClusters((prevState) => ({
      ...prevState,
      [clusterName]: !prevState[clusterName],
    }))
  }

  return (
    <Card className=''>
      <CardHeader>
        <CardTitle>K-Means Clustering Results</CardTitle>
        <CardDescription>Blogs grouped into clusters based on similarity calculated with pearsons.</CardDescription>
      </CardHeader>
      <CardContent>
        {Object.entries(kmeans).map(([clusterName, blogs]) => (
          <Collapsible key={clusterName} className='mb-4' open={openClusters[clusterName]} onOpenChange={() => handleToggle(clusterName)}>
            <div className='flex items-center justify-between px-4'>
              <h2 className='text-xl font-semibold'>{clusterName.toUpperCase()}</h2>
              <CollapsibleTrigger asChild>
                <Button variant='ghost' size='sm' className='w-9 p-0'>
                  <ChevronsUpDown className='h-4 w-4' />
                  <span className='sr-only'>Toggle</span>
                </Button>
              </CollapsibleTrigger>
            </div>
            <CollapsibleContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Blog Name</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {blogs.map((blogName, index) => (
                    <TableRow key={index}>
                      <TableCell>{blogName}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </CollapsibleContent>
          </Collapsible>
        ))}
      </CardContent>
    </Card>
  )
}
