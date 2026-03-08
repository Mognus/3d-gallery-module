export interface ModelAsset {
    id: number;
    name: string;
    model_url: string;
    canvas_mesh_name: string;
    default_scale: number;
    created_at: string;
    updated_at: string;
}

export interface GalleryImage {
    id: number;
    name: string;
    url: string;
    created_at: string;
    updated_at: string;
}

export interface GalleryItem {
    id: number;
    title: string;
    model_asset_id: number;
    model_asset: ModelAsset;
    image_id: number | null;
    image: GalleryImage | null;
    created_at: string;
    updated_at: string;
}
