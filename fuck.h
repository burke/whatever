#import <Cocoa/Cocoa.h>

extern void HandleURL(char*);

@interface FuckAppDelegate: NSObject<NSApplicationDelegate>
    - (void)handleGetURLEvent:(NSAppleEventDescriptor *) event withReplyEvent:(NSAppleEventDescriptor *)replyEvent;
@end

void RunApp();
