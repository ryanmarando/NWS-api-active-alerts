const extractYouTubeID = (url) => {
  const regex =
    /(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^\"&?\/\s]{11})/;
  const matches = url.match(regex);
  return matches ? matches[1] : null; // Returns the video ID if found
};

const YouTubeEmbed = ({ url, autoplay = false, title, className }) => {
  const videoID = extractYouTubeID(url);

  if (!videoID) {
    return <p>Invalid YouTube URL</p>; // Handle case where the video ID is invalid
  }

  const autoplayParam = autoplay ? "1" : "0";

  return (
    <div className={className}>
      {title && (
        <h2 className="text-lg font-semibold mb-2 text-center">{title}</h2>
      )}
      <div className="relative aspect-video">
        <iframe
          src={`https://www.youtube.com/embed/${videoID}?autoplay=${autoplayParam}`}
          className="absolute top-0 left-0 w-full h-full"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowFullScreen
          title={title || "YouTube Video"}
        ></iframe>
      </div>
    </div>
  );
};

export default YouTubeEmbed;
