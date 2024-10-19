import "./blog.css";
import BackgroundMedia from "@/components/ui/backgroundmedia";

export function Blog() {
    return (
        <div className="blog">
            <BackgroundMedia
                type="image"
                variant="none"
                src="/public/images/binoculars.jpg"
            ></BackgroundMedia>
        </div>
    );
}
