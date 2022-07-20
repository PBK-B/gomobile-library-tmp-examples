//
//  SwiftSupportObserve.swift
//  goMobileDemo01
//
//  Created by Bin on 2022/7/20.
//

import UIKit
import GoMobileDemoApp

class SwiftSupportObserve: NSObject, SupportSupportObserveProtocol {
    
    func onCallback(_ code: Int, msg: String?) {
        NSLog("callback code is \(code) msg is \(msg)")
    }
    
}
