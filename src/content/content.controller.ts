import { Body, Controller, Get, Post } from '@nestjs/common';
import { ContentService } from './content.service';

@Controller('content')
export class ContentController {
    constructor(
            private contentService: ContentService) {}

    @Get('list')
    async getPagesList() {
        return this.contentService.getPagesList();
    }

    @Post('')
    async getPageContent(@Body() request: { route: string }) {
        return this.contentService.getPageContent(request.route);
    }
}
