const AUDIO_TYPES = [
    "audio/midi",
    "audio/mpeg",
    "audio/mp4",
    "audio/ogg",
    "audio/x-flac",
    "audio/x-wav",
    "audio/amr",
    "audio/aac",
    "audio/x-aiff"
];

const ARCHIVE_TYPES = [
    "application/epub+zip",
    "application/zip",
    "application/x-tar",
    "application/vnd.rar",
    "application/gzip",
    "application/x-bzip2",
    "application/x-7z-compressed",
    "application/x-xz",
    "application/zstd",
];

const TABLES_TYPES = [
    "application/vnd.ms-excel",
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
    "application/vnd.oasis.opendocument.spreadsheet",
]

const POWERPOINT_TYPES = [
    "application/vnd.ms-powerpoint",
    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
    "application/vnd.oasis.opendocument.presentation",
]

const DOCUMENTS_TYPES = [
    "application/rtf",
    "application/msword",
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
    "application/vnd.oasis.opendocument.text",
];


const IMAGES_TYPES = [
    "image/jpeg",
    "image/jp2",
    "image/png",
    "image/gif",
    "image/webp",
    "image/x-canon-cr2",
    "image/tiff",
    "image/bmp",
    "image/vnd.ms-photo",
    "image/vnd.adobe.photoshop",
    "image/vnd.microsoft.icon",
    "image/heif",
    "image/vnd.dwg",
    "image/x-exr",
    "image/avif",
];

const VIDEO_TYPES = [
    "video/mp4",
    "video/x-m4v",
    "video/x-matroska",
    "video/webm",
    "video/quicktime",
    "video/x-msvideo",
    "video/x-ms-wmv",
    "video/mpeg",
    "video/x-flv",
    "video/3gpp"
];


const PDF_TYPES = ["application/pdf"]

interface IFileTypesIcons {
    [key: string]: Array<string>;
}

export const FILE_TYPES_ICONS: IFileTypesIcons = {
    "file-pdf": PDF_TYPES,
    "file-archive": ARCHIVE_TYPES,
    "file-spreadsheet": TABLES_TYPES,
    "file-powerpoint": POWERPOINT_TYPES,
    "file-word": DOCUMENTS_TYPES,
    "image": IMAGES_TYPES,
    "video": VIDEO_TYPES,
    "headphones": AUDIO_TYPES,
}
