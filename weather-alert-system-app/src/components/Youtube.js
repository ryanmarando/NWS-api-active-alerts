const extractYouTubeID = (url) => {
    const regex = /(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^\"&?\/\s]{11})/;
    const matches = url.match(regex);
    return matches ? matches[1] : null;  // Returns the video ID if found
  };


const YouTubeEmbed = ({url, width = '560', height = '315', autoplay = false }) => {
    const videoID = extractYouTubeID(url);

  if (!videoID) {
    return <p>Invalid YouTube URL</p>; // Handle case where the video ID is invalid
  }

  // Set autoplay based on the prop (1 for autoplay, 0 for no autoplay)
  const autoplayParam = autoplay ? '1' : '0';


  return (
    <div>
      <iframe
        src={`https://www.youtube.com/embed/${videoID}?autoplay=${autoplayParam}`}
        width={width}
        height={height}
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowFullScreen
        title="YouTube Video"
      ></iframe>
    </div>
  );
};

export default YouTubeEmbed;
