import { Controller, Get } from '@nestjs/common';
import { CategoryService } from './category.service';

@Controller('category')
export class CategoryController {
    constructor(
        private categoryService: CategoryService) {}

    @Get('list')
    async getCategoryList()
    {
        return await this.categoryService.getCategoryList();
    }

    @Get(':id')
    async getCategoryInfo()
    {
        return await this.categoryService.getCategoryById();
    }
}
