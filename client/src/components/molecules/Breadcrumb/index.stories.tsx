import { ComponentMeta } from '@storybook/react'
import BreadcrumbItem from 'components/atoms/BreadcrumbItem'
import Breadcrumb from './index'

export default { title: 'Molecules/Breadcrumb' } as ComponentMeta<
  typeof Breadcrumb
>

export const Standard = () => (
  <Breadcrumb>
    <BreadcrumbItem>
      <a href="#">Top</a>
    </BreadcrumbItem>
    <BreadcrumbItem>
      <a href="#">Clothes</a>
    </BreadcrumbItem>
    <BreadcrumbItem>Item</BreadcrumbItem>
  </Breadcrumb>
)
