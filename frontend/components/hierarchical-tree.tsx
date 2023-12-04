'use client'

import * as React from 'react';
import { HierarchicalNode } from '@/app/page';
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from './ui/button';

const TreeNode = ({ node, level = 0, isLast = false }: { node: HierarchicalNode, level?: number, isLast?: boolean }) => {
  const [isOpen, setIsOpen] = React.useState(true);
  const hasChildren = node.children && node.children.length > 0;
  const isRootNode = level === 0;

  const toggle = () => {
    if (hasChildren) {
      setIsOpen(!isOpen);
    }
  };

  return (
    <div className={`relative ${isRootNode ? '' : 'ml-4'}`}>
      {level !== 0 && (
        <div className={`absolute w-0.5 ${isLast ? 'h-full' : 'top-0 bottom-1/2'}`} style={{ left: '-2px' }}></div>
      )}
      <div className={`flex items-center ${hasChildren ? 'cursor-pointer' : ''}`} onClick={toggle}>
        {hasChildren && (
          <Button variant="ghost" size="sm" className="w-9 p-0 mr-2">
            <span className={`${isOpen ? 'transform rotate-90' : ''}`}>
              ➤
            </span>
          </Button>
        )}
        <span className="py-1">{node.name}</span>
      </div>
      {isOpen && hasChildren && (
      <div className="mt-1">
        {(node.children || []).map((childNode, index) => (
          <TreeNode
            key={index}
            node={childNode}
            level={level + 1}
            isLast={index === (node.children || []).length - 1}
          />
        ))}
      </div>
    )}
    </div>
  );
};

export default function HierarchicalTree({ hierarchical }: { hierarchical: HierarchicalNode[] }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Hierarchical clustering</CardTitle>
        <CardDescription>
          Interactive tree representation of hierarchical clustering.
        </CardDescription>
      </CardHeader>
      <CardContent className="overflow-auto p-4">
        {hierarchical.map((rootNode, index) => (
          <TreeNode key={index} node={rootNode} isLast={index === hierarchical.length - 1} />
        ))}
      </CardContent>
    </Card>
  );
}
