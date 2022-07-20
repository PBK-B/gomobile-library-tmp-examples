//
//  ViewController.swift
//  goMobileDemo01
//
//  Created by Bin on 2022/7/20.
//

import UIKit
import GoMobileDemoApp

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
        
        let su = SupportNewSupport()
        su!.echo()
    }

}
