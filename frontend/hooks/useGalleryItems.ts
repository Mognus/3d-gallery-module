"use client";

import useSWR from "swr";
import { fetcher } from "@/lib/api/fetcher";
import type { GalleryItem } from "../types";
import type { ListResponse } from "@/modules/admin-module/frontend/types";

export function useGalleryItems(page = 1, limit = 20) {
    return useSWR<ListResponse<GalleryItem>>(
        `/gallery/items?page=${page}&limit=${limit}`,
        fetcher,
        { revalidateOnFocus: false }
    );
}
