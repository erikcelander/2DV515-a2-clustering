"use client"

import * as React from "react";
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Table, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table";
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { ChevronsUpDown } from "lucide-react";
import { ClusterData } from '@/app/page'



export default function Clusters({ clusters }: { clusters: ClusterData }) {

  return (
    <Card className=''>
      <CardHeader>
        <CardTitle>K-Means Clustering Results</CardTitle>
        <CardDescription>
          Blogs grouped into clusters based on similarity calculated with pearsons.
        </CardDescription>
      </CardHeader>
      <CardContent>
        {Object.entries(clusters).map(([clusterName, blogs]) => (
          <Collapsible key={clusterName} className="mb-4">
            <div className="flex items-center justify-between px-4">
              <h2 className="text-xl font-semibold">{clusterName.toUpperCase()}</h2>
              <CollapsibleTrigger asChild>
                <button>
                  <ChevronsUpDown className="h-4 w-4" />
                </button>
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
  );
}
